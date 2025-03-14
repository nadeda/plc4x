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
package org.apache.plc4x.java.test.readwrite;

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

public class ComplexOfComplexUsageC implements Message {

  // Properties.
  protected final short irrelevant;

  public ComplexOfComplexUsageC(short irrelevant) {
    super();
    this.irrelevant = irrelevant;
  }

  public short getIrrelevant() {
    return irrelevant;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ComplexOfComplexUsageC");

    // Simple Field (irrelevant)
    writeSimpleField("irrelevant", irrelevant, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("ComplexOfComplexUsageC");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ComplexOfComplexUsageC _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (irrelevant)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static ComplexOfComplexUsageC staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static ComplexOfComplexUsageC staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ComplexOfComplexUsageC");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short irrelevant = readSimpleField("irrelevant", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("ComplexOfComplexUsageC");
    // Create the instance
    ComplexOfComplexUsageC _complexOfComplexUsageC;
    _complexOfComplexUsageC = new ComplexOfComplexUsageC(irrelevant);
    return _complexOfComplexUsageC;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ComplexOfComplexUsageC)) {
      return false;
    }
    ComplexOfComplexUsageC that = (ComplexOfComplexUsageC) o;
    return (getIrrelevant() == that.getIrrelevant()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getIrrelevant());
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
