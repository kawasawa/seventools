export SHELL=/bin/bash
export GO111MODULE=on


# ------------------------------------------------------------------------------
# 定数定義
# ------------------------------------------------------------------------------

# カラーコード
C_RESET  := \033[m
C_RED    := \033[31m
C_GREEN  := \033[32m
C_YELLOW := \033[33m

# ログレベル
INFO  := printf "${C_GREEN}[INFO]  ${C_RESET}"  && echo -e
WARN  := printf "${C_YELLOW}[WARN]  ${C_RESET}" && echo -e
ERROR := printf "${C_RED}[ERROR] ${C_RESET}"   && echo -e


# ------------------------------------------------------------------------------
# コマンド一覧表示
# ------------------------------------------------------------------------------

.PHONY: list
list:
	@${INFO} 'select the number of the command.';\
		echo '';\
		select SELECT_VAL in $$(cat Makefile | grep -e '.PHONY:' | grep -v 'list' | sed 's!^\.PHONY\: *!!') 'CANCEL';\
		do\
			echo '';\
			if [ "$${SELECT_VAL}" = 'CANCEL' ]; then\
				${INFO} "'CANCEL' selected. abort the process...";\
				exit 0;\
			fi;\
			if [ -z $${SELECT_VAL} ]; then\
				${WARN} 'that selection does not exist. abort the process...';\
				exit 0;\
			fi;\
			echo -e ">>> make $${SELECT_VAL}${C_RESET}";\
			make --no-print-directory "$${SELECT_VAL}";\
			break;\
		done;


# ------------------------------------------------------------------------------
# モジュール更新
# ------------------------------------------------------------------------------

.PHONY: update-module
update-module:
	@${INFO} 'go get'
	@cat go.mod | awk '/\t.+ v[0-9]+\.[0-9]+\.[a-z0-9\-\+]+$$/ { print $$1 }' | xargs -I {} go get -u -d {}
	@${INFO} 'go mod tidy'
	@go env GOVERSION | sed -r 's/^go([0-9]+\.[0-9]+).[0-9]+$$/\1/' | go mod tidy -compat=$$(cat)
	@${INFO} 'completed.'


# ------------------------------------------------------------------------------
# 静的解析
# ------------------------------------------------------------------------------

.PHONY: lint
lint:
	@${INFO} 'golangci-lint run'
	@mkdir -p ./dist; \
		golangci-lint run --out-format code-climate > ./dist/golangci-lint-report.json; \
		${INFO} 'completed.'


# ------------------------------------------------------------------------------
# ユニットテスト
# ------------------------------------------------------------------------------

.PHONY: test
test:
	@${INFO} 'go test'
	@mkdir -p ./coverage; \
	 go clean -testcache; \
	 go test ./... -coverprofile ./coverage/cover.out > ./coverage/test.out 2>&1; \
		EXIT_CODE=$$?; \
		cat ./coverage/test.out; \
		go tool cover -html=./coverage/cover.out -o ./coverage/cover.html; \
		go tool cover -func=./coverage/cover.out | grep total | tr -d '\t' | sed -e 's/(statements)/ /g'; \
		exit $$EXIT_CODE; \
	 ${INFO} 'completed.'


# ------------------------------------------------------------------------------
# ビルド
# ------------------------------------------------------------------------------

.PHONY: build
build:
	@${INFO} 'build-wasm'
	@mkdir -p ./docs; \
		GOOS=js GOARCH=wasm go build -o ./docs/main.wasm ./presentation; \
		${INFO} 'completed.'


# ------------------------------------------------------------------------------
# ローカル実行
# ------------------------------------------------------------------------------

.PHONY: exec
exec:
	@${INFO} 'exec...'
	@goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`./docs`)))'; \
		${INFO} 'completed.'
