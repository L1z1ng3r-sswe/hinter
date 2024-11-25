# Distributed System Patterns

1. [Saga Pattern](#saga-pattern)  
2. [Temporal](#temporal)  
3. [Orchestration](#orchestration)  

---

## Distributed System Patterns Details

### Saga Pattern <a id="saga-pattern"></a>

Паттерн для управления распределёнными транзакциями, который делит их на серию шагов, выполняемых последовательно или параллельно. Если какой-то шаг не удаётся, запускаются компенсирующие операции для отмены изменений предыдущих шагов.  
Пример использования: Обработка транзакций в системе электронной коммерции (create user).

### Temporal <a id="temporal"></a>

Платформа для управления рабочими процессами и распределёнными транзакциями. Temporal предоставляет автоматическое управление состоянием, ретраями и компенсирующими действиями, упрощая реализацию долгоживущих процессов.  
Преимущества:
- Автоматическое восстановление после сбоев.  
- Мониторинг состояния через веб-интерфейс.  
- Поддержка различных языков, включая Go.

### Orchestration <a id="orchestration"></a>

Модель взаимодействия микросервисов, при которой центральный оркестратор управляет выполнением шагов в процессе. Он определяет порядок выполнения операций, отслеживает состояния и обрабатывает ошибки.  
Пример: Центральный сервис управляет рабочим процессом создания нового пользователя, запуская сервисы по верификации, созданию профиля и отправке уведомлений.

### Use case:

User System: Create a user in the main database.
Profile System: Generate a user profile with default settings.
Fraud Detection System: Validate the user against fraud patterns.
Notification System: Send a welcome notification.
Email System: Send a verification email.
Logs/Analytics System: Record user creation activity for monitoring and analytics.
Compensation: Handle rollbacks if any step fails (e.g., delete user and clean up resources).


```package workflows
import (
	"go.temporal.io/sdk/workflow"
)

func CreateUserWorkflow(ctx workflow.Context, userID string, email string) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: workflow.DefaultTimeout,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// Step 1: Create User in User System
	err := workflow.ExecuteActivity(ctx, CreateUserActivity, userID).Get(ctx, nil)
	if err != nil {
		return err
	}

	// Step 2: Create Profile in Profile System
	err = workflow.ExecuteActivity(ctx, CreateProfileActivity, userID).Get(ctx, nil)
	if err != nil {
		workflow.ExecuteActivity(ctx, DeleteUserActivity, userID).Get(ctx, nil)
		return err
	}

	// Step 3: Fraud Detection
	err = workflow.ExecuteActivity(ctx, CheckFraudActivity, userID).Get(ctx, nil)
	if err != nil {
		workflow.ExecuteActivity(ctx, DeleteProfileActivity, userID).Get(ctx, nil)
		workflow.ExecuteActivity(ctx, DeleteUserActivity, userID).Get(ctx, nil)
		return err
	}

	// Step 4: Send Welcome Notification
	err = workflow.ExecuteActivity(ctx, SendNotificationActivity, userID).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Warn("Failed to send notification; continuing")
	}

	// Step 5: Send Verification Email
	err = workflow.ExecuteActivity(ctx, SendVerificationEmailActivity, email).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Warn("Failed to send email; continuing")
	}

	// Step 6: Log and Analytics
	err = workflow.ExecuteActivity(ctx, LogUserCreationActivity, userID).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Warn("Failed to log analytics; continuing")
	}

	return nil
}
```