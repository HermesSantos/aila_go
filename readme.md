# Aila – Commit Message Generator (Go)

---

# 🇺🇸 English

Aila is a Go-based utility that automatically generates commit messages using the Gemini model.  
Its goal is to standardize, streamline, and improve the quality of commit messages across Git projects.

## Features

- Automatic commit message generation using the Gemini API.
- Simple CLI, easy to integrate into any workflow.
- Automatic reading of `git diff` for contextual awareness.
- Local configuration storage using **SQLite**.
- Persistent settings for message style, API key, and per-repository preferences.

## Requirements

- Go 1.21+
- Google Gemini API key
- Git installed

## Installation

```bash
git clone https://github.com/usuario/aila
cd aila
go build -o aila
```

## Optional: Install globally

```bash
sudo mv aila /usr/local/bin/aila
sudo chmod +x /usr/local/bin/aila
```



# Aila – Commit Message Generator (Go)

---

# 🇧🇷 Português

Aila é um utilitário em Go para gerar mensagens de commit automaticamente usando o modelo Gemini.
O objetivo é padronizar, acelerar e melhorar a qualidade das mensagens de commit em projetos Git.

## Funcionalidades

- Geração automática de mensagens de commit usando a API do Gemini.
- CLI simples e fácil de integrar ao workflow.
- Leitura automática do `git diff` para contextualização.
- Armazenamento local de configurações em **SQLite**.
- Persistência de preferências de estilo, chave de API e opções específicas por repositório.

## Pré-requisitos

- Go 1.21+
- Chave de API do Google Gemini
- Git instalado

## Instalação

```bash
git clone https://github.com/usuario/aila
cd aila
go build -o aila
```

## Opcional para uso global

```bash
sudo mv aila /usr/local/bin/aila
sudo chmod +x /usr/local/bin/aila
```
