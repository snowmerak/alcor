# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/protobuf/auth.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/protobuf/auth.proto',
  package='auth',
  syntax='proto3',
  serialized_options=b'Z\nmodel/auth',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x19proto/protobuf/auth.proto\x12\x04\x61uth\"\"\n\nPrivateKey\x12\t\n\x01X\x18\x01 \x01(\t\x12\t\n\x01Y\x18\x02 \x01(\t\"!\n\tPublicKey\x12\t\n\x01X\x18\x01 \x01(\t\x12\t\n\x01Y\x18\x02 \x01(\t\"!\n\tSignature\x12\t\n\x01R\x18\x01 \x01(\t\x12\t\n\x01S\x18\x02 \x01(\tB\x0cZ\nmodel/authb\x06proto3'
)




_PRIVATEKEY = _descriptor.Descriptor(
  name='PrivateKey',
  full_name='auth.PrivateKey',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='X', full_name='auth.PrivateKey.X', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Y', full_name='auth.PrivateKey.Y', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=35,
  serialized_end=69,
)


_PUBLICKEY = _descriptor.Descriptor(
  name='PublicKey',
  full_name='auth.PublicKey',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='X', full_name='auth.PublicKey.X', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Y', full_name='auth.PublicKey.Y', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=71,
  serialized_end=104,
)


_SIGNATURE = _descriptor.Descriptor(
  name='Signature',
  full_name='auth.Signature',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='R', full_name='auth.Signature.R', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='S', full_name='auth.Signature.S', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=106,
  serialized_end=139,
)

DESCRIPTOR.message_types_by_name['PrivateKey'] = _PRIVATEKEY
DESCRIPTOR.message_types_by_name['PublicKey'] = _PUBLICKEY
DESCRIPTOR.message_types_by_name['Signature'] = _SIGNATURE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

PrivateKey = _reflection.GeneratedProtocolMessageType('PrivateKey', (_message.Message,), {
  'DESCRIPTOR' : _PRIVATEKEY,
  '__module__' : 'proto.protobuf.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.PrivateKey)
  })
_sym_db.RegisterMessage(PrivateKey)

PublicKey = _reflection.GeneratedProtocolMessageType('PublicKey', (_message.Message,), {
  'DESCRIPTOR' : _PUBLICKEY,
  '__module__' : 'proto.protobuf.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.PublicKey)
  })
_sym_db.RegisterMessage(PublicKey)

Signature = _reflection.GeneratedProtocolMessageType('Signature', (_message.Message,), {
  'DESCRIPTOR' : _SIGNATURE,
  '__module__' : 'proto.protobuf.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.Signature)
  })
_sym_db.RegisterMessage(Signature)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
