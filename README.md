# gemini-golang

CLI로 gemini에게 빠르게 뭔가 물어볼 때 사용하려고 만들었습니다.

## Setup

### linux

```shell
export GEMINI_API_KEY=<your gemini api key>
```

## Install

### Linux

```shell
git clone github.com/SlowCloud/gemini-golang
go build
```

```shell
mkdir ~/bin
mv gemini-golang ~/bin
```

## 기능

### `gemini-golang ask`

단편적인 질문과 응답이 가능합니다.

- `--long`, `-l`
  - 긴 입력을 넣을 수 있습니다.
  - 윈도우의 경우엔 `Ctrl+Z`, 리눅스는 `Ctrl+D`를 누르면 입력이 종료됩니다.

## Roadmap

- [ ] 유즈케이스 골라내기
- [ ] 설계 재진행하기
- [ ] gemini 전용 기능 활용하기

## reference

https://pkg.go.dev/google.golang.org/genai

## etc.

간단하게 만들 줄 알았는데 생각보다 추가하고 싶은 기능이 많아서? 커맨드가 그리 편한 것 같지 않아서? 더 다듬어봐야 할 것 같습니다.

개발은 설계부터 🫠
