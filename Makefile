.PHONY: up down consumer producer clean status

# 啟動 Kafka 和 Zookeeper
up:
	@echo "🚀 啟動 Kafka 和 Zookeeper..."
	docker-compose up -d
	@echo "⏳ 等待 Kafka 就緒..."
	@sleep 10
	@echo "✅ Kafka 已就緒！"

# 停止並清理容器
down:
	@echo "🛑 停止 Kafka 和 Zookeeper..."
	docker-compose down -v
	@echo "🧹 清理完成！"

# 啟動消費者
consumer:
	@echo "👂 啟動消費者..."
	cd consumer && go run main.go

# 啟動生產者
producer:
	@echo "📤 啟動生產者..."
	cd producer && go run main.go

# 清理專案資料
clean:
	@echo "🗑️  清理專案資料..."
	docker-compose down -v --remove-orphans
	@echo "✅ 清理完成！"

# 顯示狀態
status:
	@echo "📊 容器狀態："
	docker-compose ps --format table
