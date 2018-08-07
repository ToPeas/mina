package mina

import "errors"

var (
	ErrSignature   = errors.New("signature error")
	ErrPaddingSize = errors.New("padding size error")
)
