// @generated by protoc-gen-es v1.3.0
// @generated from file store/common.proto (package slash.store, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum slash.store.RowStatus
 */
export const RowStatus = proto3.makeEnum(
  "slash.store.RowStatus",
  [
    {no: 0, name: "ROW_STATUS_UNSPECIFIED"},
    {no: 1, name: "NORMAL"},
    {no: 2, name: "ARCHIVED"},
  ],
);
