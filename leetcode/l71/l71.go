package main

import (
	"fmt"
	"strings"
)

type stack struct {
	sl []string
}

func newStack() *stack {
	return &stack{
		sl: []string{},
	}
}

func (s *stack) push(str string) {
	s.sl = append(s.sl, str)
}

func (s *stack) pop() string {
	if len(s.sl) == 0 {
		return ""
	}
	ret := s.sl[len(s.sl)-1]
	s.sl = s.sl[:len(s.sl)-1]
	return ret
}

func (s *stack) peek() string {
	if len(s.sl) == 0 {
		return ""
	}
	return s.sl[len(s.sl)-1]
}

func (s *stack) size() int {
	return len(s.sl)
}
func (s *stack) isEmpty() bool {
	return len(s.sl) == 0
}

func simplifyPath(path string) string {
	s := newStack()

	i := 0
	for i < len(path) {
		if path[i] == '/' {
			j := i + 1
			for j < len(path) {
				if path[j] == '/' {
					break
				}
				j++
			}
			s.push(path[i:j])
			i = j
		}
	}

	// fmt.Println(s.sl) /////

	t := newStack()
	for !s.isEmpty() {
		switch s.peek() {
		case "/.":
			fallthrough
		case "/":
			s.pop()
			continue
		case "/..":
			popCnt := 0
			for s.peek() == "/.." && s.size() > 1 {
				popCnt++
				s.pop()
			}
			for popCnt > 0 && s.size() > 0 {
				if s.peek() == "/.." {
					popCnt++
					s.pop()
					continue
				}
				if s.peek() == "/." || s.peek() == "/" {
					s.pop()
					continue
				}
				s.pop()
				popCnt--
			}

			if s.peek() == "/.." && s.size() == 1 {
				s.pop()
			}

		default:
			t.push(s.pop())
		}
	}

	// fmt.Println(t.sl) /////

	var builder strings.Builder

	if t.isEmpty() {
		return "/"
	}

	for !t.isEmpty() {
		builder.WriteString(t.pop())
	}

	return builder.String()
}

func main() {
	// path1 := "/home/"
	// path2 := "/home//foo/"
	// path3 := "/home/user/Documents/../Pictures"
	// path4 := "/../"
	// path5 := "/.../a/../b/c/../d/./"

	// fmt.Println(simplifyPath(path1))
	// fmt.Println()
	// fmt.Println(simplifyPath(path2))
	// fmt.Println()
	// fmt.Println(simplifyPath(path3))
	// fmt.Println()
	// fmt.Println(simplifyPath(path4))
	// fmt.Println()
	// fmt.Println(simplifyPath(path5))
	// fmt.Println()

	// fmt.Println(simplifyPath("/a/./b/../../c/"))
	// fmt.Println(simplifyPath("/a//b////c/d//././/.."))
	// fmt.Println(simplifyPath("/home/../../.."))
	// fmt.Println(simplifyPath("/home/of/foo/../../bar/../../is/./here/."))
	fmt.Println(simplifyPath("///TJbrd/owxdG//"))
	fmt.Println(simplifyPath("/../lGyI/"))
}
