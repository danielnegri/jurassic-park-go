// Copyright 2023 The Jurassic Park Authors
//
// Licensed under the AGPL, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.gnu.org/licenses/agpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// OpTests composes a testsuite to run all the Ops related tests in one group
type OpTests struct {
	suite.Suite
}

func TestErrMsg(t *testing.T) {
	const op Op = "TestErrMsg"
	msg := "test error"
	err := E(op, msg)
	require.Equal(t, msg, err.Error())
}

func TestErrEmbed(t *testing.T) {
	const op Op = "TestErrEmbed"
	msg := "test error"
	childErr := errors.New(msg)
	err := E(op, childErr)
	require.Equal(t, err.Error(), childErr.Error())
}

func TestSeverity(t *testing.T) {
	const op Op = "TestSeverity"
	msg := "test error"

	err := E(op, msg)
	require.Equal(t, logrus.ErrorLevel, Severity(err))

	err = E(op, msg, logrus.WarnLevel)
	require.Equal(t, logrus.WarnLevel, Severity(err))

	err = E(op, err)
	require.Equal(t, logrus.WarnLevel, Severity(err))

	err = E(op, err, logrus.InfoLevel)
	require.Equal(t, logrus.InfoLevel, Severity(err))
}

func TestKind(t *testing.T) {
	const op Op = "TestKind"
	msg := "test error"

	err := E(op, msg, KindBadRequest)
	require.Equal(t, KindBadRequest, Kind(err))
	require.Equal(t, http.StatusText(http.StatusBadRequest), KindText(err))
}

func TestOps(t *testing.T) {
	suite.Run(t, new(OpTests))
}

func (op *OpTests) TestOps() {
	const (
		op1 Op = "TestOps.op1"
		op2 Op = "TestOps.op2"
		op3 Op = "TestOps.op3"
	)

	err1 := E(op1, "op 1")
	err2 := E(op2, err1)
	err3 := E(op3, err2)

	require.ElementsMatch(op.T(), []Op{op1, op2, op3}, Ops(err3.(Error)))
}

func (op *OpTests) SetupTest() {}

func (op *OpTests) TestString() {
	const op1 Op = "testOps.op1"
	require.Equal(op.T(), op1.String(), "testOps.op1")
}

func TestExpect(t *testing.T) {
	err := E("TestExpect", "error message", KindBadRequest)

	severity := Expect(err, KindBadRequest)
	require.Equalf(t, severity, logrus.InfoLevel, "expected an info level log but got %v", severity)

	severity = Expect(err, KindAlreadyExists)
	require.Equalf(t, severity, logrus.ErrorLevel, "expected an error level but got %v", severity)

	severity = Expect(err, KindAlreadyExists, KindBadRequest)
	require.Equalf(t, severity, logrus.InfoLevel, "expected an info level log but got %v", severity)

	severity = Expect(err, KindAlreadyExists, KindNotImplemented)
	require.Equalf(t, severity, logrus.ErrorLevel, "expected an error level but got %v", severity)
}
