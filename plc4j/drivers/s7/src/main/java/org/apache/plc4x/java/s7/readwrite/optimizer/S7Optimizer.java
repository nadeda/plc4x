/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.s7.readwrite.optimizer;

import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.s7.readwrite.*;
import org.apache.plc4x.java.s7.readwrite.context.S7DriverContext;
import org.apache.plc4x.java.s7.readwrite.tag.S7ClkTag;
import org.apache.plc4x.java.s7.readwrite.tag.S7SzlTag;
import org.apache.plc4x.java.s7.readwrite.tag.S7Tag;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.messages.DefaultPlcReadRequest;
import org.apache.plc4x.java.spi.messages.DefaultPlcWriteRequest;
import org.apache.plc4x.java.spi.messages.utils.TagValueItem;
import org.apache.plc4x.java.spi.optimizer.BaseOptimizer;

import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.LinkedList;
import java.util.List;

import org.apache.plc4x.java.s7.readwrite.tag.S7StringVarLengthTag;

public class S7Optimizer extends BaseOptimizer {

    public static final int EMPTY_READ_REQUEST_SIZE = new S7MessageRequest(0, new S7ParameterReadVarRequest(
        Collections.emptyList()), null).getLengthInBytes();
    public static final int EMPTY_READ_RESPONSE_SIZE = new S7MessageResponseData(0, new S7ParameterReadVarResponse(
        (short) 0), new S7PayloadReadVarResponse(Collections.emptyList()), (short) 0, (short) 0).getLengthInBytes();
    public static final int EMPTY_WRITE_REQUEST_SIZE = new S7MessageRequest(0, new S7ParameterWriteVarRequest(
        Collections.emptyList()), new S7PayloadWriteVarRequest(Collections.emptyList())).getLengthInBytes();
    public static final int EMPTY_WRITE_RESPONSE_SIZE = new S7MessageResponseData(0, new S7ParameterWriteVarResponse(
        (short) 0), new S7PayloadWriteVarResponse(Collections.emptyList()), (short) 0, (short) 0).getLengthInBytes();
    public static final int S7_ADDRESS_ANY_SIZE = 2 +
        new S7AddressAny(TransportSize.INT, 1, 1, MemoryArea.DATA_BLOCKS, 1, (byte) 0).getLengthInBytes();

    @Override
    protected List<PlcReadRequest> processReadRequest(PlcReadRequest readRequest, DriverContext driverContext) {
        S7DriverContext s7DriverContext = (S7DriverContext) driverContext;
        List<PlcReadRequest> processedRequests = new LinkedList<>();

        // This calculates the size of the header for the request and response.
        int curRequestSize = EMPTY_READ_REQUEST_SIZE;
        // An empty response has the same size as an empty request.
        int curResponseSize = EMPTY_READ_RESPONSE_SIZE;

        // List of all items in the current request.

        LinkedHashMap<String, PlcTag> curTags = new LinkedHashMap<>();

        for (String tagName : readRequest.getTagNames()) {
           
            //TODO: Individual processing of these types of tags. like S7StringTag
            if ((readRequest.getTag(tagName) instanceof S7SzlTag) ||
                (readRequest.getTag(tagName) instanceof S7ClkTag)) {
                curTags.put(tagName, readRequest.getTag(tagName));
                continue;
            }
            
            if ((readRequest.getTag(tagName) instanceof S7StringVarLengthTag)) {
                LinkedHashMap<String, PlcTag> strTags = new LinkedHashMap<>();
                strTags.put(tagName, readRequest.getTag(tagName));
                processedRequests.add(new DefaultPlcReadRequest(
                ((DefaultPlcReadRequest) readRequest).getReader(), strTags));
                continue;
            }            
            
            S7Tag tag = (S7Tag) readRequest.getTag(tagName);

            int readRequestItemSize = S7_ADDRESS_ANY_SIZE;
            int readResponseItemSize = 4 + (tag.getNumberOfElements() * tag.getDataType().getSizeInBytes());
            // If it's an odd number of bytes, add one to make it even
            if (readResponseItemSize % 2 == 1) {
                readResponseItemSize++;
            }

            // If adding the item would not exceed the sizes, add it to the current request.
            if (((curRequestSize + readRequestItemSize) <= s7DriverContext.getPduSize()) &&
                ((curResponseSize + readResponseItemSize) <= s7DriverContext.getPduSize())) {
                // Increase the current request sizes.
                curRequestSize += readRequestItemSize;
                curResponseSize += readResponseItemSize;

                // Add the item.
            }
            // If they would exceed, start a new request.
            else {
                // Create a new PlcReadRequest containing the current tag item.
                processedRequests.add(new DefaultPlcReadRequest(
                    ((DefaultPlcReadRequest) readRequest).getReader(), curTags));

                // Reset the size and item lists.
                curRequestSize = EMPTY_READ_REQUEST_SIZE + readRequestItemSize;
                curResponseSize = EMPTY_READ_RESPONSE_SIZE + readResponseItemSize;
                curTags = new LinkedHashMap<>();

                // Splitting of huge tags not yet implemented, throw an exception instead.
                if (((curRequestSize + readRequestItemSize) > s7DriverContext.getPduSize()) &&
                    ((curResponseSize + readResponseItemSize) > s7DriverContext.getPduSize())) {
                    throw new PlcRuntimeException("Tag size exceeds maximum payload for one item.");
                }
            }
            curTags.put(tagName, tag);
        }

        // Create a new PlcReadRequest from the remaining tag items.
        if (!curTags.isEmpty()) {
            processedRequests.add(new DefaultPlcReadRequest(
                ((DefaultPlcReadRequest) readRequest).getReader(), curTags));
        }

        return processedRequests;
    }


    @Override
    protected List<PlcWriteRequest> processWriteRequest(PlcWriteRequest writeRequest, DriverContext driverContext) {

        for (String tagName : writeRequest.getTagNames()) {
            if (writeRequest.getTag(tagName) instanceof S7ClkTag) {
                return Collections.singletonList(writeRequest);
            }
        }

        S7DriverContext s7DriverContext = (S7DriverContext) driverContext;
        List<PlcWriteRequest> processedRequests = new LinkedList<>();

        // This calculates the size of the header for the request and response.
        int curRequestSize = EMPTY_WRITE_REQUEST_SIZE;
        // An empty response has the same size as an empty request.
        int curResponseSize = EMPTY_WRITE_RESPONSE_SIZE;

        // List of all items in the current request.
        LinkedHashMap<String, TagValueItem> curTags = new LinkedHashMap<>();

        for (String tagName : writeRequest.getTagNames()) {
            
            if ((writeRequest.getTag(tagName) instanceof S7StringVarLengthTag)) {
                LinkedHashMap<String, TagValueItem> strTags = new LinkedHashMap<>();
                strTags.put(tagName, 
                        new TagValueItem(writeRequest.getTag(tagName), 
                                writeRequest.getPlcValue(tagName)));  
                processedRequests.add(new DefaultPlcWriteRequest(
                ((DefaultPlcWriteRequest) writeRequest).getWriter(), strTags));
                continue;                
            }
                                    
            S7Tag tag = (S7Tag) writeRequest.getTag(tagName);
            PlcValue value = writeRequest.getPlcValue(tagName);

            int writeRequestItemSize = S7_ADDRESS_ANY_SIZE + 4/* Size of Payload item header*/;
            if (tag.getDataType() == TransportSize.BOOL) {
                writeRequestItemSize += (int) Math.ceil((double) tag.getNumberOfElements() / 8);
            } else {
                writeRequestItemSize += (tag.getNumberOfElements() * tag.getDataType().getSizeInBytes());
            }
            // If it's an odd number of bytes, add one to make it even
            if (writeRequestItemSize % 2 == 1) {
                writeRequestItemSize++;
            }
            int writeResponseItemSize = 4;

            // If adding the item would not exceed the sizes, add it to the current request.
            if (((curRequestSize + writeRequestItemSize) <= s7DriverContext.getPduSize()) &&
                ((curResponseSize + writeResponseItemSize) <= s7DriverContext.getPduSize())) {
                // Increase the current request sizes.
                curRequestSize += writeRequestItemSize;
                curResponseSize += writeResponseItemSize;

                // Add the item.
            }
            // If adding them would exceed, start a new request.
            else {
                // Create a new PlcWriteRequest containing the current tag item.
                processedRequests.add(new DefaultPlcWriteRequest(
                    ((DefaultPlcWriteRequest) writeRequest).getWriter(), curTags));

                // Reset the size and item lists.
                curRequestSize = EMPTY_WRITE_REQUEST_SIZE + writeRequestItemSize;
                curResponseSize = EMPTY_WRITE_RESPONSE_SIZE + writeResponseItemSize;
                curTags = new LinkedHashMap<>();

                // Splitting of huge tags not yet implemented, throw an exception instead.
                if (((curRequestSize + writeRequestItemSize) > s7DriverContext.getPduSize()) &&
                    ((curResponseSize + writeResponseItemSize) > s7DriverContext.getPduSize())) {
                    throw new PlcRuntimeException("Tag size exceeds maximum payload for one item.");
                }
            }
            curTags.put(tagName, new TagValueItem(tag, value));
        }

        // Create a new PlcWriteRequest from the remaining tag items.
        if (!curTags.isEmpty()) {
            processedRequests.add(new DefaultPlcWriteRequest(
                ((DefaultPlcWriteRequest) writeRequest).getWriter(), curTags));
        }

        return processedRequests;
    }

}
