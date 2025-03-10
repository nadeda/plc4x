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
package org.apache.plc4x.java.openprotocol.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class OpenProtocolMessageResultTracesCurvePlotDataRev1
    extends OpenProtocolMessageResultTracesCurvePlotData implements Message {

  // Accessors for discriminator values.
  public Integer getRevision() {
    return (int) 1;
  }

  // Properties.
  protected final String resultDataIdentifier;
  protected final String timeStamp;
  protected final List<VariableDataField> dataFields;

  public OpenProtocolMessageResultTracesCurvePlotDataRev1(
      Integer midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber,
      String resultDataIdentifier,
      String timeStamp,
      List<VariableDataField> dataFields) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
    this.resultDataIdentifier = resultDataIdentifier;
    this.timeStamp = timeStamp;
    this.dataFields = dataFields;
  }

  public String getResultDataIdentifier() {
    return resultDataIdentifier;
  }

  public String getTimeStamp() {
    return timeStamp;
  }

  public List<VariableDataField> getDataFields() {
    return dataFields;
  }

  @Override
  protected void serializeOpenProtocolMessageResultTracesCurvePlotDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpenProtocolMessageResultTracesCurvePlotDataRev1");

    // Simple Field (resultDataIdentifier)
    writeSimpleField(
        "resultDataIdentifier",
        resultDataIdentifier,
        writeString(writeBuffer, 80),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (timeStamp)
    writeSimpleField(
        "timeStamp", timeStamp, writeString(writeBuffer, 152), WithOption.WithEncoding("ASCII"));

    // Implicit Field (numberOfParameterDataFields) (Used for parsing, but its value is not stored
    // as it's implicitly given by the objects content)
    int numberOfParameterDataFields = (int) (COUNT(getDataFields()));
    writeImplicitField(
        "numberOfParameterDataFields",
        numberOfParameterDataFields,
        writeUnsignedInt(writeBuffer, 24),
        WithOption.WithEncoding("ASCII"));

    // Array Field (dataFields)
    writeComplexTypeArrayField(
        "dataFields", dataFields, writeBuffer, WithOption.WithEncoding("ASCII"));

    writeBuffer.popContext("OpenProtocolMessageResultTracesCurvePlotDataRev1");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageResultTracesCurvePlotDataRev1 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (resultDataIdentifier)
    lengthInBits += 80;

    // Simple field (timeStamp)
    lengthInBits += 152;

    // Implicit Field (numberOfParameterDataFields)
    lengthInBits += 24;

    // Array field
    if (dataFields != null) {
      int i = 0;
      for (VariableDataField element : dataFields) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= dataFields.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static OpenProtocolMessageResultTracesCurvePlotDataBuilder
      staticParseOpenProtocolMessageResultTracesCurvePlotDataBuilder(
          ReadBuffer readBuffer, Integer revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageResultTracesCurvePlotDataRev1");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    String resultDataIdentifier =
        readSimpleField(
            "resultDataIdentifier", readString(readBuffer, 80), WithOption.WithEncoding("ASCII"));

    String timeStamp =
        readSimpleField("timeStamp", readString(readBuffer, 152), WithOption.WithEncoding("ASCII"));

    int numberOfParameterDataFields =
        readImplicitField(
            "numberOfParameterDataFields",
            readUnsignedInt(readBuffer, 24),
            WithOption.WithEncoding("ASCII"));

    List<VariableDataField> dataFields =
        readCountArrayField(
            "dataFields",
            new DataReaderComplexDefault<>(
                () -> VariableDataField.staticParse(readBuffer), readBuffer),
            numberOfParameterDataFields,
            WithOption.WithEncoding("ASCII"));

    readBuffer.closeContext("OpenProtocolMessageResultTracesCurvePlotDataRev1");
    // Create the instance
    return new OpenProtocolMessageResultTracesCurvePlotDataRev1BuilderImpl(
        resultDataIdentifier, timeStamp, dataFields);
  }

  public static class OpenProtocolMessageResultTracesCurvePlotDataRev1BuilderImpl
      implements OpenProtocolMessageResultTracesCurvePlotData
          .OpenProtocolMessageResultTracesCurvePlotDataBuilder {
    private final String resultDataIdentifier;
    private final String timeStamp;
    private final List<VariableDataField> dataFields;

    public OpenProtocolMessageResultTracesCurvePlotDataRev1BuilderImpl(
        String resultDataIdentifier, String timeStamp, List<VariableDataField> dataFields) {
      this.resultDataIdentifier = resultDataIdentifier;
      this.timeStamp = timeStamp;
      this.dataFields = dataFields;
    }

    public OpenProtocolMessageResultTracesCurvePlotDataRev1 build(
        Integer midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageResultTracesCurvePlotDataRev1
          openProtocolMessageResultTracesCurvePlotDataRev1 =
              new OpenProtocolMessageResultTracesCurvePlotDataRev1(
                  midRevision,
                  noAckFlag,
                  targetStationId,
                  targetSpindleId,
                  sequenceNumber,
                  numberOfMessageParts,
                  messagePartNumber,
                  resultDataIdentifier,
                  timeStamp,
                  dataFields);
      return openProtocolMessageResultTracesCurvePlotDataRev1;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageResultTracesCurvePlotDataRev1)) {
      return false;
    }
    OpenProtocolMessageResultTracesCurvePlotDataRev1 that =
        (OpenProtocolMessageResultTracesCurvePlotDataRev1) o;
    return (getResultDataIdentifier() == that.getResultDataIdentifier())
        && (getTimeStamp() == that.getTimeStamp())
        && (getDataFields() == that.getDataFields())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getResultDataIdentifier(), getTimeStamp(), getDataFields());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
