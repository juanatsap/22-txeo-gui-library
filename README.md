# 22-txeo-gui-library

## Descripción General
Una biblioteca de componentes de interfaz gráfica de usuario (GUI) desarrollada en Go, diseñada para facilitar la creación de aplicaciones de escritorio multiplataforma con una API coherente y fácil de usar.

## Información del Repositorio
- **Proyecto ID**: 22
- **Repositorio**: `/Users/txeo/Git/go/22-txeo-gui-library`
- **Lenguaje Principal**: Go

## Características Principales
- Framework ligero para interfaces gráficas
- Componentes reutilizables con estilo personalizable
- Sistema de layout flexible
- Manejo de eventos de usuario
- Compatibilidad multiplataforma (Windows, macOS, Linux)
- Integración con bibliotecas gráficas nativas

## Componentes Principales
- **Core**: Motor principal de renderizado y gestión de eventos
- **Widgets**: Conjunto de controles de interfaz gráfica
- **Layouts**: Sistema de organización de componentes
- **Themes**: Estilos y temas para personalizar la apariencia
- **Animations**: Sistema de animaciones y transiciones

## Estructura del Proyecto
```
22-txeo-gui-library/
├── doc/                  # Documentación
│   ├── technical/        # Documentación técnica
│   └── workflows/        # Flujos de trabajo y ejemplos
├── core/                 # Núcleo del sistema GUI
├── widgets/              # Componentes de interfaz
├── layouts/              # Gestión de layouts
├── themes/               # Sistema de temas
└── examples/             # Aplicaciones de ejemplo
```

## Instalación
```bash
go get github.com/txeo/gui-library
```

## Uso Básico
```go
package main

import (
    "github.com/txeo/gui-library/app"
    "github.com/txeo/gui-library/widgets"
)

func main() {
    // Crear aplicación
    myApp := app.New("Mi Aplicación")
    
    // Crear ventana principal
    window := app.NewWindow("Ventana Principal", 800, 600)
    
    // Crear widgets
    label := widgets.NewLabel("¡Hola, mundo!")
    button := widgets.NewButton("Haz clic aquí")
    
    // Configurar acciones
    button.OnClick(func() {
        label.SetText("¡Botón presionado!")
    })
    
    // Configurar layout
    container := widgets.NewVBox()
    container.Add(label)
    container.Add(button)
    
    // Establecer contenido de la ventana
    window.SetContent(container)
    
    // Mostrar ventana y ejecutar bucle de eventos
    window.Show()
    myApp.Run()
}
```

## Documentación Adicional
Para información más detallada sobre el diseño, implementación y uso de esta biblioteca, consulta:
- [Documentación Técnica](doc/technical/README.md)
- [Flujos de Trabajo y Ejemplos](doc/workflows/README.md)