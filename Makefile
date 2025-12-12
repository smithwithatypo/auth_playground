
.PHONY: frontend backend dev

frontend:
	cd frontend && npm run dev

backend:
	cd backend && go run main.go

dev:
	@echo "Starting frontend and backend..."
	@trap 'kill 0' EXIT; make backend & make frontend & wait
