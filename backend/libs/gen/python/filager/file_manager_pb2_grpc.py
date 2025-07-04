# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from filager import file_manager_pb2 as filager_dot_file__manager__pb2


class FileManagerServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.StartSending = channel.unary_unary(
        '/filager.FileManagerService/StartSending',
        request_serializer=filager_dot_file__manager__pb2.StartWriteRequest.SerializeToString,
        response_deserializer=filager_dot_file__manager__pb2.StartResponse.FromString,
        )
    self.SendChunk = channel.unary_unary(
        '/filager.FileManagerService/SendChunk',
        request_serializer=filager_dot_file__manager__pb2.WriteChunk.SerializeToString,
        response_deserializer=filager_dot_file__manager__pb2.WriteResponse.FromString,
        )
    self.ReadChunk = channel.unary_unary(
        '/filager.FileManagerService/ReadChunk',
        request_serializer=filager_dot_file__manager__pb2.ReadRequest.SerializeToString,
        response_deserializer=filager_dot_file__manager__pb2.GetChunk.FromString,
        )
    self.CloseSending = channel.unary_unary(
        '/filager.FileManagerService/CloseSending',
        request_serializer=filager_dot_file__manager__pb2.EndRequest.SerializeToString,
        response_deserializer=filager_dot_file__manager__pb2.EndResponse.FromString,
        )
    self.StartReading = channel.unary_unary(
        '/filager.FileManagerService/StartReading',
        request_serializer=filager_dot_file__manager__pb2.StartReadRequest.SerializeToString,
        response_deserializer=filager_dot_file__manager__pb2.StartResponse.FromString,
        )


class FileManagerServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def StartSending(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SendChunk(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ReadChunk(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CloseSending(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def StartReading(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_FileManagerServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'StartSending': grpc.unary_unary_rpc_method_handler(
          servicer.StartSending,
          request_deserializer=filager_dot_file__manager__pb2.StartWriteRequest.FromString,
          response_serializer=filager_dot_file__manager__pb2.StartResponse.SerializeToString,
      ),
      'SendChunk': grpc.unary_unary_rpc_method_handler(
          servicer.SendChunk,
          request_deserializer=filager_dot_file__manager__pb2.WriteChunk.FromString,
          response_serializer=filager_dot_file__manager__pb2.WriteResponse.SerializeToString,
      ),
      'ReadChunk': grpc.unary_unary_rpc_method_handler(
          servicer.ReadChunk,
          request_deserializer=filager_dot_file__manager__pb2.ReadRequest.FromString,
          response_serializer=filager_dot_file__manager__pb2.GetChunk.SerializeToString,
      ),
      'CloseSending': grpc.unary_unary_rpc_method_handler(
          servicer.CloseSending,
          request_deserializer=filager_dot_file__manager__pb2.EndRequest.FromString,
          response_serializer=filager_dot_file__manager__pb2.EndResponse.SerializeToString,
      ),
      'StartReading': grpc.unary_unary_rpc_method_handler(
          servicer.StartReading,
          request_deserializer=filager_dot_file__manager__pb2.StartReadRequest.FromString,
          response_serializer=filager_dot_file__manager__pb2.StartResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'filager.FileManagerService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
