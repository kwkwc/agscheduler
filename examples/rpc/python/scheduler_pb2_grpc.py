# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
import scheduler_pb2 as scheduler__pb2


class SchedulerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.AddJob = channel.unary_unary(
                '/services.Scheduler/AddJob',
                request_serializer=scheduler__pb2.Job.SerializeToString,
                response_deserializer=scheduler__pb2.Job.FromString,
                )
        self.GetJob = channel.unary_unary(
                '/services.Scheduler/GetJob',
                request_serializer=scheduler__pb2.JobId.SerializeToString,
                response_deserializer=scheduler__pb2.Job.FromString,
                )
        self.GetAllJobs = channel.unary_unary(
                '/services.Scheduler/GetAllJobs',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=scheduler__pb2.Jobs.FromString,
                )
        self.UpdateJob = channel.unary_unary(
                '/services.Scheduler/UpdateJob',
                request_serializer=scheduler__pb2.Job.SerializeToString,
                response_deserializer=scheduler__pb2.Job.FromString,
                )
        self.DeleteJob = channel.unary_unary(
                '/services.Scheduler/DeleteJob',
                request_serializer=scheduler__pb2.JobId.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.DeleteAllJobs = channel.unary_unary(
                '/services.Scheduler/DeleteAllJobs',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.PauseJob = channel.unary_unary(
                '/services.Scheduler/PauseJob',
                request_serializer=scheduler__pb2.JobId.SerializeToString,
                response_deserializer=scheduler__pb2.Job.FromString,
                )
        self.ResumeJob = channel.unary_unary(
                '/services.Scheduler/ResumeJob',
                request_serializer=scheduler__pb2.JobId.SerializeToString,
                response_deserializer=scheduler__pb2.Job.FromString,
                )
        self.RunJob = channel.unary_unary(
                '/services.Scheduler/RunJob',
                request_serializer=scheduler__pb2.Job.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.ScheduleJob = channel.unary_unary(
                '/services.Scheduler/ScheduleJob',
                request_serializer=scheduler__pb2.Job.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.Start = channel.unary_unary(
                '/services.Scheduler/Start',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.Stop = channel.unary_unary(
                '/services.Scheduler/Stop',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )


class SchedulerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def AddJob(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetJob(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllJobs(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateJob(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteJob(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteAllJobs(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PauseJob(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ResumeJob(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RunJob(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ScheduleJob(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Start(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Stop(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SchedulerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'AddJob': grpc.unary_unary_rpc_method_handler(
                    servicer.AddJob,
                    request_deserializer=scheduler__pb2.Job.FromString,
                    response_serializer=scheduler__pb2.Job.SerializeToString,
            ),
            'GetJob': grpc.unary_unary_rpc_method_handler(
                    servicer.GetJob,
                    request_deserializer=scheduler__pb2.JobId.FromString,
                    response_serializer=scheduler__pb2.Job.SerializeToString,
            ),
            'GetAllJobs': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAllJobs,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=scheduler__pb2.Jobs.SerializeToString,
            ),
            'UpdateJob': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateJob,
                    request_deserializer=scheduler__pb2.Job.FromString,
                    response_serializer=scheduler__pb2.Job.SerializeToString,
            ),
            'DeleteJob': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteJob,
                    request_deserializer=scheduler__pb2.JobId.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'DeleteAllJobs': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteAllJobs,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'PauseJob': grpc.unary_unary_rpc_method_handler(
                    servicer.PauseJob,
                    request_deserializer=scheduler__pb2.JobId.FromString,
                    response_serializer=scheduler__pb2.Job.SerializeToString,
            ),
            'ResumeJob': grpc.unary_unary_rpc_method_handler(
                    servicer.ResumeJob,
                    request_deserializer=scheduler__pb2.JobId.FromString,
                    response_serializer=scheduler__pb2.Job.SerializeToString,
            ),
            'RunJob': grpc.unary_unary_rpc_method_handler(
                    servicer.RunJob,
                    request_deserializer=scheduler__pb2.Job.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'ScheduleJob': grpc.unary_unary_rpc_method_handler(
                    servicer.ScheduleJob,
                    request_deserializer=scheduler__pb2.Job.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'Start': grpc.unary_unary_rpc_method_handler(
                    servicer.Start,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'Stop': grpc.unary_unary_rpc_method_handler(
                    servicer.Stop,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'services.Scheduler', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Scheduler(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def AddJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/AddJob',
            scheduler__pb2.Job.SerializeToString,
            scheduler__pb2.Job.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/GetJob',
            scheduler__pb2.JobId.SerializeToString,
            scheduler__pb2.Job.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAllJobs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/GetAllJobs',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            scheduler__pb2.Jobs.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/UpdateJob',
            scheduler__pb2.Job.SerializeToString,
            scheduler__pb2.Job.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/DeleteJob',
            scheduler__pb2.JobId.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteAllJobs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/DeleteAllJobs',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PauseJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/PauseJob',
            scheduler__pb2.JobId.SerializeToString,
            scheduler__pb2.Job.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ResumeJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/ResumeJob',
            scheduler__pb2.JobId.SerializeToString,
            scheduler__pb2.Job.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RunJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/RunJob',
            scheduler__pb2.Job.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ScheduleJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/ScheduleJob',
            scheduler__pb2.Job.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Start(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/Start',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Stop(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/services.Scheduler/Stop',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
