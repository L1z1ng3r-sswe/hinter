### 1. **Подготовка кода**
   - Код уже есть в репозитории Git.
   - Есть Dockerfile для сборки образа.

---

### 2. **Сборка Docker-образа**
   - **GitLab CI/CD** триггерится на коммит в ветки `main` или `develop`.
   - В пайплайне выполняется:
     1. `docker build` — сборка образа.
     2. `docker tag` — тегирование образа (по коммиту и latest).
     3. `docker push` — пуш образа в **Google Container Registry (GCR)**.

---

### 3. **Деплой на Kubernetes**
   - Для **develop**:
     - Обновляется staging-среда: 
       ```bash
       kubectl set image deployment/staging app=gcr.io/project/image:commit-sha
       ```
     - Проверяется статус раскатки:
       ```bash
       kubectl rollout status deployment/staging
       ```
   - Для **main**:
     - Обновляется production-среда:
       ```bash
       kubectl set image deployment/production app=gcr.io/project/image:commit-sha
       ```
     - Проверяется статус раскатки:
       ```bash
       kubectl rollout status deployment/production
       ```

---

### 4. **Управление ветками**
   - **develop** → Деплой в **staging**.
   - **main** → Деплой в **production**.

---

### Итог
**Цепочка:**  
1. **Коммит в ветку** (`main` или `develop`).  
2. GitLab CI/CD:
   - Сборка Docker-образа.
   - Пуш в GCR.
   - Деплой на Kubernetes.  
3. **Готово!**

**Важно:** Разделение окружений и использование правильных Kubernetes deployment.