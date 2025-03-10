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

public class Node extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "260";
  }

  // Properties.
  protected final NodeId nodeId;
  protected final NodeClass nodeClass;
  protected final QualifiedName browseName;
  protected final LocalizedText displayName;
  protected final LocalizedText description;
  protected final long writeMask;
  protected final long userWriteMask;
  protected final int noOfReferences;
  protected final List<ExtensionObjectDefinition> references;

  public Node(
      NodeId nodeId,
      NodeClass nodeClass,
      QualifiedName browseName,
      LocalizedText displayName,
      LocalizedText description,
      long writeMask,
      long userWriteMask,
      int noOfReferences,
      List<ExtensionObjectDefinition> references) {
    super();
    this.nodeId = nodeId;
    this.nodeClass = nodeClass;
    this.browseName = browseName;
    this.displayName = displayName;
    this.description = description;
    this.writeMask = writeMask;
    this.userWriteMask = userWriteMask;
    this.noOfReferences = noOfReferences;
    this.references = references;
  }

  public NodeId getNodeId() {
    return nodeId;
  }

  public NodeClass getNodeClass() {
    return nodeClass;
  }

  public QualifiedName getBrowseName() {
    return browseName;
  }

  public LocalizedText getDisplayName() {
    return displayName;
  }

  public LocalizedText getDescription() {
    return description;
  }

  public long getWriteMask() {
    return writeMask;
  }

  public long getUserWriteMask() {
    return userWriteMask;
  }

  public int getNoOfReferences() {
    return noOfReferences;
  }

  public List<ExtensionObjectDefinition> getReferences() {
    return references;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("Node");

    // Simple Field (nodeId)
    writeSimpleField("nodeId", nodeId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (nodeClass)
    writeSimpleEnumField(
        "nodeClass",
        "NodeClass",
        nodeClass,
        new DataWriterEnumDefault<>(
            NodeClass::getValue, NodeClass::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (browseName)
    writeSimpleField("browseName", browseName, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (displayName)
    writeSimpleField("displayName", displayName, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (writeMask)
    writeSimpleField("writeMask", writeMask, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (userWriteMask)
    writeSimpleField("userWriteMask", userWriteMask, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (noOfReferences)
    writeSimpleField("noOfReferences", noOfReferences, writeSignedInt(writeBuffer, 32));

    // Array Field (references)
    writeComplexTypeArrayField("references", references, writeBuffer);

    writeBuffer.popContext("Node");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    Node _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nodeId)
    lengthInBits += nodeId.getLengthInBits();

    // Simple field (nodeClass)
    lengthInBits += 32;

    // Simple field (browseName)
    lengthInBits += browseName.getLengthInBits();

    // Simple field (displayName)
    lengthInBits += displayName.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (writeMask)
    lengthInBits += 32;

    // Simple field (userWriteMask)
    lengthInBits += 32;

    // Simple field (noOfReferences)
    lengthInBits += 32;

    // Array field
    if (references != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : references) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= references.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("Node");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId nodeId =
        readSimpleField(
            "nodeId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    NodeClass nodeClass =
        readEnumField(
            "nodeClass",
            "NodeClass",
            new DataReaderEnumDefault<>(NodeClass::enumForValue, readUnsignedLong(readBuffer, 32)));

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

    LocalizedText description =
        readSimpleField(
            "description",
            new DataReaderComplexDefault<>(
                () -> LocalizedText.staticParse(readBuffer), readBuffer));

    long writeMask = readSimpleField("writeMask", readUnsignedLong(readBuffer, 32));

    long userWriteMask = readSimpleField("userWriteMask", readUnsignedLong(readBuffer, 32));

    int noOfReferences = readSimpleField("noOfReferences", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> references =
        readCountArrayField(
            "references",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("287")),
                readBuffer),
            noOfReferences);

    readBuffer.closeContext("Node");
    // Create the instance
    return new NodeBuilderImpl(
        nodeId,
        nodeClass,
        browseName,
        displayName,
        description,
        writeMask,
        userWriteMask,
        noOfReferences,
        references);
  }

  public static class NodeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId nodeId;
    private final NodeClass nodeClass;
    private final QualifiedName browseName;
    private final LocalizedText displayName;
    private final LocalizedText description;
    private final long writeMask;
    private final long userWriteMask;
    private final int noOfReferences;
    private final List<ExtensionObjectDefinition> references;

    public NodeBuilderImpl(
        NodeId nodeId,
        NodeClass nodeClass,
        QualifiedName browseName,
        LocalizedText displayName,
        LocalizedText description,
        long writeMask,
        long userWriteMask,
        int noOfReferences,
        List<ExtensionObjectDefinition> references) {
      this.nodeId = nodeId;
      this.nodeClass = nodeClass;
      this.browseName = browseName;
      this.displayName = displayName;
      this.description = description;
      this.writeMask = writeMask;
      this.userWriteMask = userWriteMask;
      this.noOfReferences = noOfReferences;
      this.references = references;
    }

    public Node build() {
      Node node =
          new Node(
              nodeId,
              nodeClass,
              browseName,
              displayName,
              description,
              writeMask,
              userWriteMask,
              noOfReferences,
              references);
      return node;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Node)) {
      return false;
    }
    Node that = (Node) o;
    return (getNodeId() == that.getNodeId())
        && (getNodeClass() == that.getNodeClass())
        && (getBrowseName() == that.getBrowseName())
        && (getDisplayName() == that.getDisplayName())
        && (getDescription() == that.getDescription())
        && (getWriteMask() == that.getWriteMask())
        && (getUserWriteMask() == that.getUserWriteMask())
        && (getNoOfReferences() == that.getNoOfReferences())
        && (getReferences() == that.getReferences())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getNodeId(),
        getNodeClass(),
        getBrowseName(),
        getDisplayName(),
        getDescription(),
        getWriteMask(),
        getUserWriteMask(),
        getNoOfReferences(),
        getReferences());
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
