# Makefile برای Go URL Shortener

APP_NAME = go-url-shortener
MAIN = cmd/main.go

# استفاده از Go default
GO = go

# Default target
.PHONY: all
all: run

# اجرا کردن برنامه
.PHONY: run
run:
	-$(GO) run $(MAIN)

# Build باینری
.PHONY: build
build:
	$(GO) build -o $(APP_NAME) $(MAIN)

# اجرای تست‌ها
.PHONY: test
test:
	$(GO) test ./... -v

# مرتب‌سازی و به‌روز کردن ماژول‌ها
.PHONY: tidy
tidy:
	$(GO) mod tidy

# پاکسازی باینری
.PHONY: clean
clean:
	rm -f $(APP_NAME)
