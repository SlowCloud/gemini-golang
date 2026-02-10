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

## Features

- Gemini API 활용한 채팅 기능
- 채팅 저장 기능
    - `~/histories`에 저장됩니다.
- 저장된 채팅을 선택하여 이어서 채팅을 진행하는 기능

## Roadmap

- Gemini 모델 선택 기능
- Gemini 이외의 AI 모델 활용 기능

## reference

https://pkg.go.dev/google.golang.org/genai
