# gemini-golang

터미널 환경에서 Gemini AI와 채팅할 수 있는 도구입니다.

Gemini CLI와 비교했을 때, 최소 기능만이 구현되어 있어
훨씬 가볍고 빠르게 채팅을 시작하고 진행할 수 있습니다.

## Setup

사용을 위해, 환경 변수에 API Key를 등록하는 과정이 필요합니다.

AI Key는 [Google AI Studio](https://aistudio.google.com/api-keys)에서
발급이 가능합니다.

### Linux

```bash
export GOOGLE_API_KEY=YOUR_GEMINI_API_KEY
```

### Windows

```powershell
$env:GOOGLE_API_KEY=YOUR_GEMINI_API_KEY
```

## Install

```shell
go install github.com/SlowCloud/gemini-golang@latest
```

## reference

https://pkg.go.dev/google.golang.org/genai
