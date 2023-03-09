# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: deliver.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rdeliver.proto\"\x1d\n\rSimpleRequest\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\t\"\xb7\x02\n\x07\x42\x61rData\x12\r\n\x05Insid\x18\x01 \x01(\t\x12\x0f\n\x07Ts_open\x18\x02 \x01(\x03\x12\x12\n\nOpen_price\x18\x03 \x01(\x02\x12\x12\n\nHigh_price\x18\x04 \x01(\x02\x12\x11\n\tLow_price\x18\x05 \x01(\x02\x12\x13\n\x0b\x43lose_price\x18\x06 \x01(\x02\x12\x0b\n\x03Vol\x18\x07 \x01(\x02\x12\x0e\n\x06VolCcy\x18\x08 \x01(\x02\x12\x13\n\x0bVolCcyQuote\x18\t \x01(\x02\x12\n\n\x02Oi\x18\n \x01(\x02\x12\r\n\x05OiCcy\x18\x0b \x01(\x02\x12\r\n\x05Ts_oi\x18\x0c \x01(\x03\x12\x13\n\x0b\x46undingRate\x18\r \x01(\x02\x12\x17\n\x0fNextFundingRate\x18\x0e \x01(\x02\x12\x16\n\x0eTs_FundingRate\x18\x0f \x01(\x03\x12\x1a\n\x12TS_NextFundingRate\x18\x10 \x01(\x03\"}\n\x08TickData\x12\r\n\x05Insid\x18\x01 \x01(\t\x12\x10\n\x08Ts_Price\x18\x02 \x01(\x03\x12\x12\n\nAsk1_price\x18\x03 \x01(\x02\x12\x12\n\nBid1_price\x18\x04 \x01(\x02\x12\x13\n\x0b\x41sk1_volumn\x18\x05 \x01(\x02\x12\x13\n\x0b\x42id1_volumn\x18\x06 \x01(\x02\"\x1f\n\x08Response\x12\x13\n\x0bresponse_me\x18\x01 \x01(\t\"\x1c\n\x08JsonInfo\x12\x10\n\x08jsoninfo\x18\x01 \x01(\t\"\xb6\x01\n\x0bLocalSubmit\x12\x0f\n\x07subtype\x18\x01 \x01(\t\x12\x11\n\tbarcustom\x18\x02 \x01(\t\x12\x11\n\ttickInsid\x18\x03 \x01(\t\x12\x10\n\x08\x62\x61rInsid\x18\x04 \x01(\t\x12\x10\n\x08tickPort\x18\x05 \x01(\t\x12\x0f\n\x07\x62\x61rPort\x18\x06 \x01(\t\x12\x13\n\x0b\x61\x63\x63ountPort\x18\x07 \x01(\t\x12\x14\n\x0cstrategyname\x18\x08 \x01(\t\x12\x10\n\x08initjson\x18\t \x01(\t\"\xf5\x02\n\x05Order\x12\r\n\x05insId\x18\x01 \x01(\t\x12\x0e\n\x06tdMode\x18\x02 \x01(\t\x12\x0b\n\x03\x63\x63y\x18\x03 \x01(\t\x12\x0f\n\x07\x63lOrdId\x18\x04 \x01(\t\x12\x0b\n\x03tag\x18\x05 \x01(\t\x12\x0c\n\x04side\x18\x06 \x01(\t\x12\x0f\n\x07posSide\x18\x07 \x01(\t\x12\x0f\n\x07ordType\x18\x08 \x01(\t\x12\n\n\x02sz\x18\t \x01(\t\x12\n\n\x02px\x18\n \x01(\t\x12\x12\n\nreduceOnly\x18\x0b \x01(\x08\x12\x0e\n\x06tgtCcy\x18\x0c \x01(\t\x12\x10\n\x08\x62\x61nAmend\x18\r \x01(\x08\x12\x13\n\x0btpTriggerPx\x18\x0e \x01(\t\x12\x0f\n\x07tpOrdPx\x18\x0f \x01(\t\x12\x13\n\x0bslTriggerPx\x18\x10 \x01(\t\x12\x0f\n\x07slOrdPx\x18\x11 \x01(\t\x12\x17\n\x0ftpTriggerPxType\x18\x12 \x01(\t\x12\x17\n\x0fslTriggerPxType\x18\x13 \x01(\t\x12\x14\n\x0cquickMgnType\x18\x14 \x01(\t\x12\x10\n\x08\x62rokerID\x18\x15 \x01(\t2A\n\x12\x43ustomDataReceiver\x12+\n\x12\x43ustomDataReceiver\x12\x08.BarData\x1a\t.Response\"\x00\x32I\n\x14SubmitServerReceiver\x12\x31\n\x14SubmitServerReceiver\x12\x0c.LocalSubmit\x1a\t.Response\"\x00\x32>\n\x10TickDataReceiver\x12*\n\x10TickDataReceiver\x12\t.TickData\x1a\t.Response\"\x00\x32\x34\n\x0cOrerReceiver\x12$\n\rOrerRReceiver\x12\x06.Order\x1a\t.Response\"\x00\x32\x36\n\x0cJsonReceiver\x12&\n\x0cJsonReceiver\x12\t.JsonInfo\x1a\t.Response\"\x00\x42\x0bZ\t.;deliverb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'deliver_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\t.;deliver'
  _SIMPLEREQUEST._serialized_start=17
  _SIMPLEREQUEST._serialized_end=46
  _BARDATA._serialized_start=49
  _BARDATA._serialized_end=360
  _TICKDATA._serialized_start=362
  _TICKDATA._serialized_end=487
  _RESPONSE._serialized_start=489
  _RESPONSE._serialized_end=520
  _JSONINFO._serialized_start=522
  _JSONINFO._serialized_end=550
  _LOCALSUBMIT._serialized_start=553
  _LOCALSUBMIT._serialized_end=735
  _ORDER._serialized_start=738
  _ORDER._serialized_end=1111
  _CUSTOMDATARECEIVER._serialized_start=1113
  _CUSTOMDATARECEIVER._serialized_end=1178
  _SUBMITSERVERRECEIVER._serialized_start=1180
  _SUBMITSERVERRECEIVER._serialized_end=1253
  _TICKDATARECEIVER._serialized_start=1255
  _TICKDATARECEIVER._serialized_end=1317
  _ORERRECEIVER._serialized_start=1319
  _ORERRECEIVER._serialized_end=1371
  _JSONRECEIVER._serialized_start=1373
  _JSONRECEIVER._serialized_end=1427
# @@protoc_insertion_point(module_scope)
