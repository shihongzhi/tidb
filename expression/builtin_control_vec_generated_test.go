// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go generate in expression/generator; DO NOT EDIT.

package expression

import (
	"testing"

	. "github.com/pingcap/check"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/tidb/types"
)

var vecBuiltinControlCases = map[string][]vecExprBenchCase{
	ast.If: {

		{types.ETInt, []types.EvalType{types.ETInt, types.ETInt, types.ETInt}, nil},

		{types.ETReal, []types.EvalType{types.ETInt, types.ETReal, types.ETReal}, nil},

		{types.ETDecimal, []types.EvalType{types.ETInt, types.ETDecimal, types.ETDecimal}, nil},

		{types.ETString, []types.EvalType{types.ETInt, types.ETString, types.ETString}, nil},

		{types.ETDatetime, []types.EvalType{types.ETInt, types.ETDatetime, types.ETDatetime}, nil},

		{types.ETDuration, []types.EvalType{types.ETInt, types.ETDuration, types.ETDuration}, nil},

		{types.ETJson, []types.EvalType{types.ETInt, types.ETJson, types.ETJson}, nil},
	},
}

func (s *testEvaluatorSuite) TestVectorizedBuiltinControlEvalOneVec(c *C) {
	testVectorizedEvalOneVec(c, vecBuiltinControlCases)
}

func (s *testEvaluatorSuite) TestVectorizedBuiltinControlFunc(c *C) {
	testVectorizedBuiltinFunc(c, vecBuiltinControlCases)
}

func BenchmarkVectorizedBuiltinControlEvalOneVec(b *testing.B) {
	benchmarkVectorizedEvalOneVec(b, vecBuiltinControlCases)
}

func BenchmarkVectorizedBuiltinControlFunc(b *testing.B) {
	benchmarkVectorizedBuiltinFunc(b, vecBuiltinControlCases)
}
