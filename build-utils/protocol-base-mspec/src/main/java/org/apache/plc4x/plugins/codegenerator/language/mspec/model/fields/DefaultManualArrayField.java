/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 */

package org.apache.plc4x.plugins.codegenerator.language.mspec.model.fields;

import org.apache.plc4x.plugins.codegenerator.types.fields.ManualArrayField;
import org.apache.plc4x.plugins.codegenerator.types.references.TypeReference;
import org.apache.plc4x.plugins.codegenerator.types.terms.Term;

public class DefaultManualArrayField extends DefaultTaggedField implements ManualArrayField {

    private final TypeReference type;
    private final String name;
    private final LoopType loopType;
    private final Term loopExpression;
    private final Term serializeExpression;
    private final Term parseExpression;
    private final Term lengthExpression;
    private final Term[] params;

    public DefaultManualArrayField(String[] tags, TypeReference type, String name, LoopType loopType, Term loopExpression, Term serializeExpression, Term parseExpression, Term lengthExpression, Term[] params) {
        super(tags);
        this.type = type;
        this.name = name;
        this.loopType = loopType;
        this.loopExpression = loopExpression;
        this.serializeExpression = serializeExpression;
        this.parseExpression = parseExpression;
        this.lengthExpression = lengthExpression;
        this.params = params;
    }

    public TypeReference getType() {
        return type;
    }

    public String getName() {
        return name;
    }

    public LoopType getLoopType() {
        return loopType;
    }

    public Term getLoopExpression() {
        return loopExpression;
    }

    public Term getSerializeExpression() {
        return serializeExpression;
    }

    public Term getParseExpression() {
        return parseExpression;
    }

    public Term getLengthExpression() {
        return lengthExpression;
    }

    @Override
    public Term[] getParams() {
        return params;
    }

}
