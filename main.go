go get github.com/tealeg/xlsx
go get gonum.org/v1/gonum/mat
package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
    "gonum.org/v1/gonum/mat"
    "gonum.org/v1/gonum/optimize"
    "math/rand"
    "time"
)

func main() {
    // Configurar la semilla para reproducibilidad
    rand.Seed(time.Now().UnixNano())

    // Cargar el archivo XLSX
    xlFile, err := xlsx.OpenFile("Catalogo1960_2021.xlsx")
    if err != nil {
        fmt.Println("Error al abrir el archivo XLSX:", err)
        return
    }

    // Obtener la hoja deseada (en este caso, "Sheet1")
    sheetName := "Sheet1"
    sheet, ok := xlFile.Sheet[sheetName]
    if !ok {
        fmt.Println("No se encontró la hoja:", sheetName)
        return
    }

    // Preparar datos y etiquetas
    data, labels := prepareData(sheet.Rows[1:]) // Se omite la fila de encabezado

    // Crear un modelo de red neuronal
    inputSize := 6      // Número de características de entrada
    hiddenSize := 10    // Tamaño de la capa oculta
    outputSize := 1     // Número de salidas
    learningRate := 0.1 // Tasa de aprendizaje

    model := createNeuralNetwork(inputSize, hiddenSize, outputSize)

    // Entrenar el modelo
    optimizer := optimize.NewRMSProp(learningRate, 0.9, 1e-6, -1, 1000, nil)
    params := optimize.Params{
        GradTol: 1e-5,
        FeatTol: 1e-5,
        Func:    model,
    }
    x, _ := optimizer.Minimize(params, nil, data, labels)

    // Imprimir los parámetros entrenados
    fmt.Println("Parámetros entrenados:")
    fmt.Println(x)

    // Utilizar el modelo entrenado para hacer predicciones (no se muestra aquí)
    // ...
}

func prepareData(rows []*xlsx.Row) (*mat.Dense, *mat.VecDense) {
    numRows := len(rows)
    numFeatures := 6 // Número de características (ajusta según tus datos reales)

    data := mat.NewDense(numRows, numFeatures, nil)
    labels := mat.NewVecDense(numRows, nil)

    for i, row := range rows {
        for j := 0; j < numFeatures; j++ {
            // Reemplaza con la extracción de datos reales de las celdas
            // Puedes usar row.Cells[j] para acceder a las celdas de cada columna
            data.Set(i, j, rand.Float64())
        }
        // Reemplaza con la obtención de la etiqueta real
        // labels.SetVec(i, valor_real_de_etiqueta)
        labels.SetVec(i, rand.Float64())
    }

    return data, labels
}

func createNeuralNetwork(inputSize, hiddenSize, outputSize int) *optimize.Func {
    // Definir la arquitectura de la red neuronal y la función de costo
    // Implementar la propagación hacia adelante y hacia atrás, el cálculo del costo y los gradientes
    // Devolver una función de costo personalizada
    return nil
}
