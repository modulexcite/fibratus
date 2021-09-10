/*
 * Copyright 2019-2020 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kparams

const (
	// Null is a null parameter type
	Null Type = iota
	// UnicodeString a string of 16-bit characters. By default, assumed to have been encoded using UTF-16LE
	UnicodeString
	// AnsiString a string of 8-bit characters
	AnsiString
	// Int8 a signed 8-bit integer
	Int8
	// Uint8 an unsigned 8-bit integer
	Uint8
	// Int16 a signed 16-bit integer
	Int16
	// Uint16 an unsigned 16-bit integer
	Uint16
	// Int32 a signed 32-bit integer
	Int32
	// Uint32 an unsigned 32-bit integer
	Uint32
	// Int64 a signed 64-bit integer
	Int64
	// Uint64 an unsigned 64-bit integer
	Uint64
	// Float an IEEE 4-byte floating-point number
	Float
	// Double an IEEE 8-byte floating-point number
	Double
	// Bool a 32-bit value where 0 is false and 1 is true
	Bool
	// Binary is a binary data of variable size. The size must be specified in the data definition as a constant or a reference to another (integer) data item.For an IP V6 address, the data should be an IN6_ADDR structure.
	// For a socket address, the data should be a SOCKADDR_STORAGE structure. The AF_INET, AF_INET6, and AF_LINK address families are supported
	Binary
	// Pointer an unsigned 32-bit or 64-bit pointer value. The size depends on the architecture of the computer logging the event
	Pointer
	// PID is the process identifier
	PID
	// TID is the thread identifier
	TID
	// HexInt8 is the hexadecimal representation of 8-bit integer
	HexInt8
	// HexInt16 is the hexadecimal representation of 16-bit integer
	HexInt16
	// HexInt32 is the hexadecimal representation of 32-bit integer
	HexInt32
	// HexInt64 is the hexadecimal representation of 64-bit integer
	HexInt64
	// Port represents the endpoint port number
	Port
	// IP is the IP address
	IP
	// IPv4 is the IPv4 address
	IPv4
	// IPv6 is the IPv6 address
	IPv6
	// Time represents the timestamp
	Time
	// Slice represents a collection of items
	Slice
	// Enum represents an enumeration
	Enum
	// Map represents a map
	Map
	// Object is the generic object
	Object
	// Unknown represent an unknown parameter type
	Unknown
)
