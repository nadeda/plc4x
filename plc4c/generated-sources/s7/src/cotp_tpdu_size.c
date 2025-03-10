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

#include <string.h>
#include <plc4c/driver_s7_static.h>

#include "cotp_tpdu_size.h"

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_s7_read_write_cotp_tpdu_size plc4c_s7_read_write_cotp_tpdu_size_null_const;

plc4c_s7_read_write_cotp_tpdu_size plc4c_s7_read_write_cotp_tpdu_size_null() {
  return plc4c_s7_read_write_cotp_tpdu_size_null_const;
}

// Parse function.
plc4c_return_code plc4c_s7_read_write_cotp_tpdu_size_parse(plc4x_spi_context ctx, plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_cotp_tpdu_size* _message) {
    plc4c_return_code _res = OK;

    uint8_t value;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
    *_message = plc4c_s7_read_write_cotp_tpdu_size_for_value(value);

    return _res;
}

plc4c_return_code plc4c_s7_read_write_cotp_tpdu_size_serialize(plc4x_spi_context ctx, plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_cotp_tpdu_size* _message) {
    plc4c_return_code _res = OK;

    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, *_message);

    return _res;
}

plc4c_s7_read_write_cotp_tpdu_size plc4c_s7_read_write_cotp_tpdu_size_for_value(uint8_t value) {
    for(int i = 0; i < plc4c_s7_read_write_cotp_tpdu_size_num_values(); i++) {
        if(plc4c_s7_read_write_cotp_tpdu_size_value_for_index(i) == value) {
            return plc4c_s7_read_write_cotp_tpdu_size_value_for_index(i);
        }
    }
    return -1;
}

plc4c_s7_read_write_cotp_tpdu_size plc4c_s7_read_write_cotp_tpdu_size_value_of(char* value_string) {
    if(strcmp(value_string, "SIZE_128") == 0) {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_128;
    }
    if(strcmp(value_string, "SIZE_256") == 0) {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_256;
    }
    if(strcmp(value_string, "SIZE_512") == 0) {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_512;
    }
    if(strcmp(value_string, "SIZE_1024") == 0) {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_1024;
    }
    if(strcmp(value_string, "SIZE_2048") == 0) {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_2048;
    }
    if(strcmp(value_string, "SIZE_4096") == 0) {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_4096;
    }
    if(strcmp(value_string, "SIZE_8192") == 0) {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_8192;
    }
    return -1;
}

int plc4c_s7_read_write_cotp_tpdu_size_num_values() {
  return 7;
}

plc4c_s7_read_write_cotp_tpdu_size plc4c_s7_read_write_cotp_tpdu_size_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_128;
      }
      case 1: {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_256;
      }
      case 2: {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_512;
      }
      case 3: {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_1024;
      }
      case 4: {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_2048;
      }
      case 5: {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_4096;
      }
      case 6: {
        return plc4c_s7_read_write_cotp_tpdu_size_SIZE_8192;
      }
      default: {
        return -1;
      }
    }
}

uint16_t plc4c_s7_read_write_cotp_tpdu_size_get_size_in_bytes(plc4c_s7_read_write_cotp_tpdu_size value) {
  switch(value) {
    case plc4c_s7_read_write_cotp_tpdu_size_SIZE_128: { /* '0x07' */
      return 128;
    }
    case plc4c_s7_read_write_cotp_tpdu_size_SIZE_256: { /* '0x08' */
      return 256;
    }
    case plc4c_s7_read_write_cotp_tpdu_size_SIZE_512: { /* '0x09' */
      return 512;
    }
    case plc4c_s7_read_write_cotp_tpdu_size_SIZE_1024: { /* '0x0a' */
      return 1024;
    }
    case plc4c_s7_read_write_cotp_tpdu_size_SIZE_2048: { /* '0x0b' */
      return 2048;
    }
    case plc4c_s7_read_write_cotp_tpdu_size_SIZE_4096: { /* '0x0c' */
      return 4096;
    }
    case plc4c_s7_read_write_cotp_tpdu_size_SIZE_8192: { /* '0x0d' */
      return 8192;
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_cotp_tpdu_size plc4c_s7_read_write_cotp_tpdu_size_get_first_enum_for_field_size_in_bytes(uint16_t value) {
    switch(value) {
        case 1024: {
            return plc4c_s7_read_write_cotp_tpdu_size_SIZE_1024;
        }
        case 128: {
            return plc4c_s7_read_write_cotp_tpdu_size_SIZE_128;
        }
        case 2048: {
            return plc4c_s7_read_write_cotp_tpdu_size_SIZE_2048;
        }
        case 256: {
            return plc4c_s7_read_write_cotp_tpdu_size_SIZE_256;
        }
        case 4096: {
            return plc4c_s7_read_write_cotp_tpdu_size_SIZE_4096;
        }
        case 512: {
            return plc4c_s7_read_write_cotp_tpdu_size_SIZE_512;
        }
        case 8192: {
            return plc4c_s7_read_write_cotp_tpdu_size_SIZE_8192;
        }
        default: {
            return -1;
        }
    }
}

uint16_t plc4c_s7_read_write_cotp_tpdu_size_length_in_bytes(plc4x_spi_context ctx, plc4c_s7_read_write_cotp_tpdu_size* _message) {
    return plc4c_s7_read_write_cotp_tpdu_size_length_in_bits(ctx, _message) / 8;
}

uint16_t plc4c_s7_read_write_cotp_tpdu_size_length_in_bits(plc4x_spi_context ctx, plc4c_s7_read_write_cotp_tpdu_size* _message) {
    return 8;
}
