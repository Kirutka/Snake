# 🐍 Змейка на Go [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org/)

Простая реализация классической игры "Змейка" на языке Go с использованием библиотеки Ebiten.
 
<p align="center">
  <img src="https://github.com/user-attachments/assets/392a626a-c031-4bd9-b9c7-f302f95bbb5d" alt="Game Preview" width="600">
</p>

---

## 🚀 Как запустить

### Требования
- Установленный [Go](https://golang.org/dl/) (версия 1.16 или выше).
- Установленная библиотека [Ebiten](https://ebiten.org/).

### Установка и запуск
1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/Kirutka/Lagman.git
   cd Змейка
   ```

2. Установите зависимости:
   ```bash
   go mod download
   ```

3. Запустите игру:
   ```bash
   go run main.go
   ```

---

## 🎮 Управление
- **W** — движение вверх.
- **S** — движение вниз.
- **A** — движение влево.
- **D** — движение вправо.
- **R** — перезапуск игры после завершения.

---

## 🛠️ Особенности реализации
- **Змейка** состоит из сегментов, которые увеличиваются при съедании еды.
- **Еда** появляется в случайном месте на игровом поле.
- Игра завершается, если змейка сталкивается с границами экрана или сама с собой.
- Реализована возможность перезапуска игры после завершения.

---

## 📂 Структура проекта
```
.
├── main.go          # Основной файл игры
├── go.mod           # Файл зависимостей Go
├── go.sum           # Файл контрольных сумм зависимостей
└── README.md        # Этот файл
```

---

## 🛠️ Технологии
- **[Go](https://golang.org/)** — язык программирования.
- **[Ebiten](https://ebiten.org/)** — библиотека для создания 2D-игр.


---

## 📞 Контакты
Если у вас есть вопросы или предложения, свяжитесь со мной:  
[![Telegram](https://img.shields.io/badge/-@Lesnoy_umorust-2CA5E0?style=flat&logo=telegram)](https://t.me/Lesnoy_umorust)
[![VK](https://img.shields.io/badge/-VK-0077FF?style=flat&logo=vk)](https://vk.com/id549536760)
[![Email](https://img.shields.io/badge/-kirillzaporozec48@gmail.com-D14836?style=flat&logo=gmail)](mailto:kirillzaporozec48@gmail.com)

---

**Приятной игры!** 🎮
