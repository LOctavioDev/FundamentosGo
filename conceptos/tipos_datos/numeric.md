# Tipos de Datos Numéricos en Go

## Enteros

### `int`
- **Descripción:** Tipo de entero de tamaño dependiente de la arquitectura del sistema (32 o 64 bits).
- **Rango:** Depende de la arquitectura del sistema (32 o 64 bits).

### `int8`
- **Descripción:** Entero con signo de 8 bits.
- **Rango:** -128 a 127.

### `int16`
- **Descripción:** Entero con signo de 16 bits.
- **Rango:** -32,768 a 32,767.

### `int32`
- **Descripción:** Entero con signo de 32 bits.
- **Rango:** -2,147,483,648 a 2,147,483,647.

### `int64`
- **Descripción:** Entero con signo de 64 bits.
- **Rango:** -9,223,372,036,854,775,808 a 9,223,372,036,854,775,807.

## Enteros sin signo

### `uint`
- **Descripción:** Tipo de entero sin signo de tamaño dependiente de la arquitectura del sistema (32 o 64 bits).
- **Rango:** Depende de la arquitectura del sistema (32 o 64 bits).

### `uint8` (alias `byte`)
- **Descripción:** Entero sin signo de 8 bits.
- **Rango:** 0 a 255.

### `uint16`
- **Descripción:** Entero sin signo de 16 bits.
- **Rango:** 0 a 65,535.

### `uint32`
- **Descripción:** Entero sin signo de 32 bits.
- **Rango:** 0 a 4,294,967,295.

### `uint64`
- **Descripción:** Entero sin signo de 64 bits.
- **Rango:** 0 a 18,446,744,073,709,551,615.

### `uintptr`
- **Descripción:** Entero sin signo lo suficientemente grande para contener el valor de un puntero.
- **Rango:** Depende de la arquitectura del sistema (32 o 64 bits).

## Números de punto flotante

### `float32`
- **Descripción:** Número de punto flotante de 32 bits.
- **Rango:** Aproximadamente ±1.18e-38 a ±3.4e38.
- **Precisión:** 6-9 dígitos decimales.

### `float64`
- **Descripción:** Número de punto flotante de 64 bits.
- **Rango:** Aproximadamente ±2.23e-308 a ±1.8e308.
- **Precisión:** 15-17 dígitos decimales.

## Números complejos

### `complex64`
- **Descripción:** Número complejo con partes reales e imaginarias de 32 bits.
- **Componentes:** `float32` real y `float32` imaginario.

### `complex128`
- **Descripción:** Número complejo con partes reales e imaginarias de 64 bits.
- **Componentes:** `float64` real y `float64` imaginario.

## `rune`
- **Descripción:** Alias para `int32`, utilizado para representar un punto de código Unicode.
- **Rango:** -2,147,483,648 a 2,147,483,647.
