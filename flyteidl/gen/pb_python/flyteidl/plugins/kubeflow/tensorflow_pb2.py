# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/plugins/kubeflow/tensorflow.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import tasks_pb2 as flyteidl_dot_core_dot_tasks__pb2
from flyteidl.plugins.kubeflow import common_pb2 as flyteidl_dot_plugins_dot_kubeflow_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n*flyteidl/plugins/kubeflow/tensorflow.proto\x12\x19\x66lyteidl.plugins.kubeflow\x1a\x19\x66lyteidl/core/tasks.proto\x1a&flyteidl/plugins/kubeflow/common.proto\"\xa8\x03\n!DistributedTensorflowTrainingTask\x12l\n\x0fworker_replicas\x18\x01 \x01(\x0b\x32\x43.flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpecR\x0eworkerReplicas\x12\x64\n\x0bps_replicas\x18\x02 \x01(\x0b\x32\x43.flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpecR\npsReplicas\x12j\n\x0e\x63hief_replicas\x18\x03 \x01(\x0b\x32\x43.flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpecR\rchiefReplicas\x12\x43\n\nrun_policy\x18\x04 \x01(\x0b\x32$.flyteidl.plugins.kubeflow.RunPolicyR\trunPolicy\"\xe5\x01\n(DistributedTensorflowTrainingReplicaSpec\x12\x1a\n\x08replicas\x18\x01 \x01(\x05R\x08replicas\x12\x14\n\x05image\x18\x02 \x01(\tR\x05image\x12\x36\n\tresources\x18\x03 \x01(\x0b\x32\x18.flyteidl.core.ResourcesR\tresources\x12O\n\x0erestart_policy\x18\x04 \x01(\x0e\x32(.flyteidl.plugins.kubeflow.RestartPolicyR\rrestartPolicyB\xef\x01\n\x1d\x63om.flyteidl.plugins.kubeflowB\x0fTensorflowProtoP\x01Z7github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/plugins\xa2\x02\x03\x46PK\xaa\x02\x19\x46lyteidl.Plugins.Kubeflow\xca\x02\x19\x46lyteidl\\Plugins\\Kubeflow\xe2\x02%Flyteidl\\Plugins\\Kubeflow\\GPBMetadata\xea\x02\x1b\x46lyteidl::Plugins::Kubeflowb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.plugins.kubeflow.tensorflow_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\035com.flyteidl.plugins.kubeflowB\017TensorflowProtoP\001Z7github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/plugins\242\002\003FPK\252\002\031Flyteidl.Plugins.Kubeflow\312\002\031Flyteidl\\Plugins\\Kubeflow\342\002%Flyteidl\\Plugins\\Kubeflow\\GPBMetadata\352\002\033Flyteidl::Plugins::Kubeflow'
  _globals['_DISTRIBUTEDTENSORFLOWTRAININGTASK']._serialized_start=141
  _globals['_DISTRIBUTEDTENSORFLOWTRAININGTASK']._serialized_end=565
  _globals['_DISTRIBUTEDTENSORFLOWTRAININGREPLICASPEC']._serialized_start=568
  _globals['_DISTRIBUTEDTENSORFLOWTRAININGREPLICASPEC']._serialized_end=797
# @@protoc_insertion_point(module_scope)
