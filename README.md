# SistemasDistribuidos

## Integrantes
Joaquin Calderon 201973571-3

## Asignación de las máquinas
Las máquinas utilizadas corresponden a dist145, dist146, dist147 y dist148. Estas fueron distribuidas de la siguiente forma:

- dist145: Central y Laboratorio Renca
- dist146: Laboratorio Pohang
- dist147: Laboratorio Kampala
- dist148: Laboratorio Pripiat

## Ejecución

Se deberá conectar a las máquinas mencionadas anteriormente mediante ssh 5 veces, 1 para la Central y 4 para los Laboratorios. Luego en donde corresponda se deberá ejecutar los códigos para Central o Laboratorio.

### Central
Primero se deberá activar la central.

Para ejecutar el código de la central se deberá ejecutar la siguiente línea en la máquina:
```
make central
```
### Laboratorio
Cuando ya esté ejecutandose la central, iniciaremos los laboratorios.

Para ejecutar el código de los Laboratorios se deberá ejecutar la siguiente línea en cada máquina:
```
make laboratorio
```
