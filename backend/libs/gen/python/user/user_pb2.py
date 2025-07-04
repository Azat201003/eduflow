# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user/user.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='user/user.proto',
  package='user',
  syntax='proto3',
  serialized_pb=_b('\n\x0fuser/user.proto\x12\x04user\"G\n\rFilterRequest\x12\x1c\n\x06\x66ilter\x18\x01 \x01(\x0b\x32\x0c.user.Filter\x12\x18\n\x04page\x18\x02 \x01(\x0b\x32\n.user.Page\"$\n\x04Page\x12\x0c\n\x04size\x18\x01 \x01(\r\x12\x0e\n\x06number\x18\x02 \x01(\r\"D\n\x06\x46ilter\x12\x10\n\x08\x63ontains\x18\x01 \x03(\t\x12\x0e\n\x06\x61\x64mins\x18\x02 \x03(\x04\x12\x18\n\x04type\x18\x03 \x01(\x0e\x32\n.user.Type\"2\n\x0c\x43reditionals\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"\x10\n\x02Id\x12\n\n\x02id\x18\x01 \x01(\x04\"H\n\x04User\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x10\n\x08username\x18\x02 \x01(\t\x12\x10\n\x08is_staff\x18\x04 \x01(\x08\x12\x10\n\x08\x66olowers\x18\x05 \x03(\x04\"\x16\n\x05Token\x12\r\n\x05token\x18\x01 \x01(\t\"\x97\x01\n\x05Group\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x11\n\twhitelist\x18\x04 \x03(\x04\x12\x11\n\tblacklist\x18\x05 \x03(\x04\x12\x0f\n\x07members\x18\x06 \x03(\x04\x12\x0e\n\x06\x61\x64mins\x18\x07 \x03(\x04\x12\x18\n\x04type\x18\x08 \x01(\x0e\x32\n.user.Type*\x1f\n\x04Type\x12\n\n\x06PUBLIC\x10\x00\x12\x0b\n\x07PRIVATE\x10\x01\x32\x8f\x02\n\x0bUserService\x12+\n\x08Register\x12\x12.user.Creditionals\x1a\x0b.user.Token\x12(\n\x05Login\x12\x12.user.Creditionals\x1a\x0b.user.Token\x12\x1f\n\x07GetUser\x12\x08.user.Id\x1a\n.user.User\x12)\n\x0eGetUserByToken\x12\x0b.user.Token\x1a\n.user.User\x12\x36\n\x11GetFilteredGroups\x12\x13.user.FilterRequest\x1a\n.user.User0\x01\x12%\n\x0cGetGroupById\x12\x08.user.Id\x1a\x0b.user.GroupB\rZ\x0bgen/go/userb\x06proto3')
)

_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='user.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='PUBLIC', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PRIVATE', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=528,
  serialized_end=559,
)
_sym_db.RegisterEnumDescriptor(_TYPE)

Type = enum_type_wrapper.EnumTypeWrapper(_TYPE)
PUBLIC = 0
PRIVATE = 1



_FILTERREQUEST = _descriptor.Descriptor(
  name='FilterRequest',
  full_name='user.FilterRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='filter', full_name='user.FilterRequest.filter', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page', full_name='user.FilterRequest.page', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=25,
  serialized_end=96,
)


_PAGE = _descriptor.Descriptor(
  name='Page',
  full_name='user.Page',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='size', full_name='user.Page.size', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='number', full_name='user.Page.number', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=98,
  serialized_end=134,
)


_FILTER = _descriptor.Descriptor(
  name='Filter',
  full_name='user.Filter',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='contains', full_name='user.Filter.contains', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='admins', full_name='user.Filter.admins', index=1,
      number=2, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='user.Filter.type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=136,
  serialized_end=204,
)


_CREDITIONALS = _descriptor.Descriptor(
  name='Creditionals',
  full_name='user.Creditionals',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='username', full_name='user.Creditionals.username', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='password', full_name='user.Creditionals.password', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=206,
  serialized_end=256,
)


_ID = _descriptor.Descriptor(
  name='Id',
  full_name='user.Id',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='user.Id.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=258,
  serialized_end=274,
)


_USER = _descriptor.Descriptor(
  name='User',
  full_name='user.User',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='user.User.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='username', full_name='user.User.username', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_staff', full_name='user.User.is_staff', index=2,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='folowers', full_name='user.User.folowers', index=3,
      number=5, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=276,
  serialized_end=348,
)


_TOKEN = _descriptor.Descriptor(
  name='Token',
  full_name='user.Token',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='token', full_name='user.Token.token', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=350,
  serialized_end=372,
)


_GROUP = _descriptor.Descriptor(
  name='Group',
  full_name='user.Group',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='user.Group.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='user.Group.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='user.Group.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='whitelist', full_name='user.Group.whitelist', index=3,
      number=4, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='blacklist', full_name='user.Group.blacklist', index=4,
      number=5, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='members', full_name='user.Group.members', index=5,
      number=6, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='admins', full_name='user.Group.admins', index=6,
      number=7, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='user.Group.type', index=7,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=375,
  serialized_end=526,
)

_FILTERREQUEST.fields_by_name['filter'].message_type = _FILTER
_FILTERREQUEST.fields_by_name['page'].message_type = _PAGE
_FILTER.fields_by_name['type'].enum_type = _TYPE
_GROUP.fields_by_name['type'].enum_type = _TYPE
DESCRIPTOR.message_types_by_name['FilterRequest'] = _FILTERREQUEST
DESCRIPTOR.message_types_by_name['Page'] = _PAGE
DESCRIPTOR.message_types_by_name['Filter'] = _FILTER
DESCRIPTOR.message_types_by_name['Creditionals'] = _CREDITIONALS
DESCRIPTOR.message_types_by_name['Id'] = _ID
DESCRIPTOR.message_types_by_name['User'] = _USER
DESCRIPTOR.message_types_by_name['Token'] = _TOKEN
DESCRIPTOR.message_types_by_name['Group'] = _GROUP
DESCRIPTOR.enum_types_by_name['Type'] = _TYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FilterRequest = _reflection.GeneratedProtocolMessageType('FilterRequest', (_message.Message,), dict(
  DESCRIPTOR = _FILTERREQUEST,
  __module__ = 'user.user_pb2'
  # @@protoc_insertion_point(class_scope:user.FilterRequest)
  ))
_sym_db.RegisterMessage(FilterRequest)

Page = _reflection.GeneratedProtocolMessageType('Page', (_message.Message,), dict(
  DESCRIPTOR = _PAGE,
  __module__ = 'user.user_pb2'
  # @@protoc_insertion_point(class_scope:user.Page)
  ))
_sym_db.RegisterMessage(Page)

Filter = _reflection.GeneratedProtocolMessageType('Filter', (_message.Message,), dict(
  DESCRIPTOR = _FILTER,
  __module__ = 'user.user_pb2'
  # @@protoc_insertion_point(class_scope:user.Filter)
  ))
_sym_db.RegisterMessage(Filter)

Creditionals = _reflection.GeneratedProtocolMessageType('Creditionals', (_message.Message,), dict(
  DESCRIPTOR = _CREDITIONALS,
  __module__ = 'user.user_pb2'
  # @@protoc_insertion_point(class_scope:user.Creditionals)
  ))
_sym_db.RegisterMessage(Creditionals)

Id = _reflection.GeneratedProtocolMessageType('Id', (_message.Message,), dict(
  DESCRIPTOR = _ID,
  __module__ = 'user.user_pb2'
  # @@protoc_insertion_point(class_scope:user.Id)
  ))
_sym_db.RegisterMessage(Id)

User = _reflection.GeneratedProtocolMessageType('User', (_message.Message,), dict(
  DESCRIPTOR = _USER,
  __module__ = 'user.user_pb2'
  # @@protoc_insertion_point(class_scope:user.User)
  ))
_sym_db.RegisterMessage(User)

Token = _reflection.GeneratedProtocolMessageType('Token', (_message.Message,), dict(
  DESCRIPTOR = _TOKEN,
  __module__ = 'user.user_pb2'
  # @@protoc_insertion_point(class_scope:user.Token)
  ))
_sym_db.RegisterMessage(Token)

Group = _reflection.GeneratedProtocolMessageType('Group', (_message.Message,), dict(
  DESCRIPTOR = _GROUP,
  __module__ = 'user.user_pb2'
  # @@protoc_insertion_point(class_scope:user.Group)
  ))
_sym_db.RegisterMessage(Group)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\013gen/go/user'))

_USERSERVICE = _descriptor.ServiceDescriptor(
  name='UserService',
  full_name='user.UserService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=562,
  serialized_end=833,
  methods=[
  _descriptor.MethodDescriptor(
    name='Register',
    full_name='user.UserService.Register',
    index=0,
    containing_service=None,
    input_type=_CREDITIONALS,
    output_type=_TOKEN,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Login',
    full_name='user.UserService.Login',
    index=1,
    containing_service=None,
    input_type=_CREDITIONALS,
    output_type=_TOKEN,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetUser',
    full_name='user.UserService.GetUser',
    index=2,
    containing_service=None,
    input_type=_ID,
    output_type=_USER,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetUserByToken',
    full_name='user.UserService.GetUserByToken',
    index=3,
    containing_service=None,
    input_type=_TOKEN,
    output_type=_USER,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetFilteredGroups',
    full_name='user.UserService.GetFilteredGroups',
    index=4,
    containing_service=None,
    input_type=_FILTERREQUEST,
    output_type=_USER,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetGroupById',
    full_name='user.UserService.GetGroupById',
    index=5,
    containing_service=None,
    input_type=_ID,
    output_type=_GROUP,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_USERSERVICE)

DESCRIPTOR.services_by_name['UserService'] = _USERSERVICE

# @@protoc_insertion_point(module_scope)
