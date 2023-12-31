package cerrors

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"runtime"
)

type UserName string
type Op string
type Kind uint8
type ServiceCode string
type ServiceMessage string

var Separator = ":\n\t"

const (
	Other      Kind = iota // 분류되지 않은 오류의 경우 (이 값은 메시지에 포함하지 않음).
	Invalid                // 유효하지 않은 행위를 한 경우.
	Auth                   // 비인증.
	Permission             // 권한이 옳바르지 않은 경우.
	IO                     // I/O에 문제가 있는 경우 (네트워트 오류 등).
	Exist                  // 이미 존재하는 경우.
	NotExist               // 존재하지 않는 경우.
	Internal               // 로직 오류의 경우.
)

type Error struct {
	User           UserName       // 요청자
	Op             Op             // 도메인/액션
	Kind           Kind           // 에러 종류
	Err            error          // 에러
	ServiceCode    ServiceCode    // 클라이언트 전용 코드
	ServiceMessage ServiceMessage // 클라이언트 전용 메시지
}

// 문자열을 이용한 에러 생성
type errorString struct {
	s string
}

// Error 관련
func (e *Error) Error() string {
	b := new(bytes.Buffer)
	if e.Op != "" {
		pad(b, ": ")
		b.WriteString(string(e.Op))
	}
	if e.User != "" {
		pad(b, ", ")
		b.WriteString("user ")
		b.WriteString(string(e.User))
	}
	if e.Kind != 0 {
		pad(b, ": ")
		b.WriteString(e.Kind.String())
	}
	if e.Err != nil {
		var prevErr *Error
		if errors.As(e.Err, &prevErr) {
			if !prevErr.isZero() {
				pad(b, Separator)
				b.WriteString(e.Err.Error())
			}
		} else {
			pad(b, ": ")
			b.WriteString(e.Err.Error())
		}
	}
	if b.Len() == 0 {
		return "no error"
	}
	return b.String()
}

func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to cerrors.E with no arguments")
	}
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case UserName:
			e.User = arg
		case Op:
			e.Op = arg
		case string:
			e.Err = Str(arg)
		case Kind:
			e.Kind = arg
		case *Error:
			copy := *arg
			e.Err = &copy
		case error:
			e.Err = arg
		case ServiceCode:
			e.ServiceCode = arg
		case ServiceMessage:
			e.ServiceMessage = arg
		default:
			_, file, line, _ := runtime.Caller(1)
			log.Printf("cerrors.E: bad call from %s:%d: %v", file, line, args)
			return Errorf("unknown type %T, value %v in error call", arg, arg)
		}
	}

	var prev *Error
	ok := errors.As(e.Err, &prev)
	if !ok {
		return e
	}

	// 중첩 에러의 경우 중복된 경우 중복 제거
	if prev.User == e.User {
		prev.User = ""
	}
	if prev.Kind == e.Kind {
		prev.Kind = Other
	}
	// 중첩 에러 일때 현재 발생한 에러가 처리하지 않은 에러라면 에러를 끌어올림
	if e.Kind == Other {
		e.Kind = prev.Kind
		prev.Kind = Other
	}
	return e
}

func (k Kind) String() string {
	switch k {
	case Other:
		return "other error"
	case Invalid:
		return "invalid operation"
	case Auth:
		return "unauthorized"
	case Permission:
		return "permission denied"
	case IO:
		return "I/O error"
	case Exist:
		return "item already exists"
	case NotExist:
		return "item does not exist"
	case Internal:
		return "internal error"
	}
	return "unknown error kind"
}

func (e *Error) isZero() bool {
	return e.User == "" && e.Op == "" && e.Kind == 0 && e.Err == nil
}

func pad(b *bytes.Buffer, str string) {
	if b.Len() == 0 {
		return
	}
	b.WriteString(str)
}

func Match(err1, err2 error) bool {
	var e1 *Error
	ok := errors.As(err1, &e1)
	if !ok {
		return false
	}
	var e2 *Error
	ok = errors.As(err2, &e2)
	if !ok {
		return false
	}
	if e1.User != "" && e2.User != e1.User {
		return false
	}
	if e1.Op != "" && e2.Op != e1.Op {
		return false
	}
	if e1.Kind != Other && e2.Kind != e1.Kind {
		return false
	}
	if e1.Err != nil {
		var cErr *Error
		if errors.As(e1.Err, &cErr) {
			return Match(e1.Err, e2.Err)
		}
		if e2.Err == nil || e2.Err.Error() != e1.Err.Error() {
			return false
		}
	}
	return true
}

func Is(kind Kind, err error) bool {
	var e *Error
	ok := errors.As(err, &e)
	if !ok {
		return false
	}
	if e.Kind != Other {
		return e.Kind == kind
	}
	if e.Err != nil {
		return Is(kind, e.Err)
	}
	return false
}

// Str errorString 관련
func Str(text string) error {
	return &errorString{text}
}

func (e *errorString) Error() string {
	return e.s
}

func Errorf(format string, args ...interface{}) error {
	return &errorString{fmt.Sprintf(format, args...)}
}
