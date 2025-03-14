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
package org.apache.plc4x.java.opcua.readwrite;

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

public class ReferenceDescription extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "520";
  }

  // Properties.
  protected final NodeId referenceTypeId;
  protected final boolean isForward;
  protected final ExpandedNodeId nodeId;
  protected final QualifiedName browseName;
  protected final LocalizedText displayName;
  protected final NodeClass nodeClass;
  protected final ExpandedNodeId typeDefinition;

  public ReferenceDescription(
      NodeId referenceTypeId,
      boolean isForward,
      ExpandedNodeId nodeId,
      QualifiedName browseName,
      LocalizedText displayName,
      NodeClass nodeClass,
      ExpandedNodeId typeDefinition) {
    super();
    this.referenceTypeId = referenceTypeId;
    this.isForward = isForward;
    this.nodeId = nodeId;
    this.browseName = browseName;
    this.displayName = displayName;
    this.nodeClass = nodeClass;
    this.typeDefinition = typeDefinition;
  }

  public NodeId getReferenceTypeId() {
    return referenceTypeId;
  }

  public boolean getIsForward() {
    return isForward;
  }

  public ExpandedNodeId getNodeId() {
    return nodeId;
  }

  public QualifiedName getBrowseName() {
    return browseName;
  }

  public LocalizedText getDisplayName() {
    return displayName;
  }

  public NodeClass getNodeClass() {
    return nodeClass;
  }

  public ExpandedNodeId getTypeDefinition() {
    return typeDefinition;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ReferenceDescription");

    // Simple Field (referenceTypeId)
    writeSimpleField(
        "referenceTypeId", referenceTypeId, new DataWriterComplexDefault<>(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (isForward)
    writeSimpleField("isForward", isForward, writeBoolean(writeBuffer));

    // Simple Field (nodeId)
    writeSimpleField("nodeId", nodeId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (browseName)
    writeSimpleField("browseName", browseName, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (displayName)
    writeSimpleField("displayName", displayName, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (nodeClass)
    writeSimpleEnumField(
        "nodeClass",
        "NodeClass",
        nodeClass,
        new DataWriterEnumDefault<>(
            NodeClass::getValue, NodeClass::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (typeDefinition)
    writeSimpleField("typeDefinition", typeDefinition, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ReferenceDescription");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ReferenceDescription _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (referenceTypeId)
    lengthInBits += referenceTypeId.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (isForward)
    lengthInBits += 1;

    // Simple field (nodeId)
    lengthInBits += nodeId.getLengthInBits();

    // Simple field (browseName)
    lengthInBits += browseName.getLengthInBits();

    // Simple field (displayName)
    lengthInBits += displayName.getLengthInBits();

    // Simple field (nodeClass)
    lengthInBits += 32;

    // Simple field (typeDefinition)
    lengthInBits += typeDefinition.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ReferenceDescription");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId referenceTypeId =
        readSimpleField(
            "referenceTypeId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean isForward = readSimpleField("isForward", readBoolean(readBuffer));

    ExpandedNodeId nodeId =
        readSimpleField(
            "nodeId",
            new DataReaderComplexDefault<>(
                () -> ExpandedNodeId.staticParse(readBuffer), readBuffer));

    QualifiedName browseName =
        readSimpleField(
            "browseName",
            new DataReaderComplexDefault<>(
                () -> QualifiedName.staticParse(readBuffer), readBuffer));

    LocalizedText displayName =
        readSimpleField(
            "displayName",
            new DataReaderComplexDefault<>(
                () -> LocalizedText.staticParse(readBuffer), readBuffer));

    NodeClass nodeClass =
        readEnumField(
            "nodeClass",
            "NodeClass",
            new DataReaderEnumDefault<>(NodeClass::enumForValue, readUnsignedLong(readBuffer, 32)));

    ExpandedNodeId typeDefinition =
        readSimpleField(
            "typeDefinition",
            new DataReaderComplexDefault<>(
                () -> ExpandedNodeId.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("ReferenceDescription");
    // Create the instance
    return new ReferenceDescriptionBuilderImpl(
        referenceTypeId, isForward, nodeId, browseName, displayName, nodeClass, typeDefinition);
  }

  public static class ReferenceDescriptionBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId referenceTypeId;
    private final boolean isForward;
    private final ExpandedNodeId nodeId;
    private final QualifiedName browseName;
    private final LocalizedText displayName;
    private final NodeClass nodeClass;
    private final ExpandedNodeId typeDefinition;

    public ReferenceDescriptionBuilderImpl(
        NodeId referenceTypeId,
        boolean isForward,
        ExpandedNodeId nodeId,
        QualifiedName browseName,
        LocalizedText displayName,
        NodeClass nodeClass,
        ExpandedNodeId typeDefinition) {
      this.referenceTypeId = referenceTypeId;
      this.isForward = isForward;
      this.nodeId = nodeId;
      this.browseName = browseName;
      this.displayName = displayName;
      this.nodeClass = nodeClass;
      this.typeDefinition = typeDefinition;
    }

    public ReferenceDescription build() {
      ReferenceDescription referenceDescription =
          new ReferenceDescription(
              referenceTypeId,
              isForward,
              nodeId,
              browseName,
              displayName,
              nodeClass,
              typeDefinition);
      return referenceDescription;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ReferenceDescription)) {
      return false;
    }
    ReferenceDescription that = (ReferenceDescription) o;
    return (getReferenceTypeId() == that.getReferenceTypeId())
        && (getIsForward() == that.getIsForward())
        && (getNodeId() == that.getNodeId())
        && (getBrowseName() == that.getBrowseName())
        && (getDisplayName() == that.getDisplayName())
        && (getNodeClass() == that.getNodeClass())
        && (getTypeDefinition() == that.getTypeDefinition())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getReferenceTypeId(),
        getIsForward(),
        getNodeId(),
        getBrowseName(),
        getDisplayName(),
        getNodeClass(),
        getTypeDefinition());
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
