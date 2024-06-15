## Notifications Lib

Библиотека для отправки уведомлений из сервисов mrm в nats. <br/>

### Установка

```sh
    go get https://gitlab.com/etagi/mrm/notifications_lib
```

### Примеры

- Установка и закрытие соединения
    ```go
        nc := notifications_lib.NewNotificationClient("nats://127.0.0.1:4222")
        defer nc.Close()   
    ```
- Инициализация проекта
    ```go
        nc := notifications_lib.NewNotificationClient("nats://127.0.0.1:4222")
  
        project := model.NewProject("Библиотека", "lib", false, []model.Type{
            {
                Name:       "test lib",
                Code:       "test lib",
                Switchable: false,
                Disabled:   false,
            },
            {
                Name:       "test lib 2",
                Code:       "test lib 2",
                Switchable: true,
                Disabled:   false,
            },
        })
  
        err := nc.ProjectPublish(project)
        if err != nil {
           log.Fatalf(err)
        }
    ```
- Отправка сообщения
    ```go
        nc := notifications_lib.NewNotificationClient("nats://127.0.0.1:4222")

        actionCreate := model.NewNotificationsAction("Create", "primary", "#")
        actionDelete := model.NewNotificationsAction("Delete", "secondary", "#")
        var actions []model.NotificationsAction
        actions = append(actions, *actionCreate)
        actions = append(actions, *actionDelete)

        notifications := model.NewNotificationsInput("content", "remove_collection", []string{"111"}, "test lib", "test lib from other service", time.Now(), actions)

        err := nc.NotificationPublish(notifications)
        if err != nil {
            log.Fatalf(err)
        }
    ```
  
Все модели можно инициализировать через конструктор или напрямую
```go
    pj := model.Project{
        Name:     "Example",
        Code:     "example",
        Disabled: false,
        Types:    []model.Type{
            {
                Name:       "type example",
                Code:       "type example",
                Switchable: false,
                Disabled:   false,
            },
        },
    }
    // or
    typeCreate := model.NewType("name", "code", false, false)
    types := []model.Type
    types = append(types, typeCreate)
	
    pj := model.NewProject("Example", "example", false, types)
```
<br/>

### Кнопки

Доступны два вида кнопок по цветам:
```go
    actionCreate := model.NewNotificationsAction("Create", "primary", "#")
    actionDelete := model.NewNotificationsAction("Delete", "secondary", "#")
```

