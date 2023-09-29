package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Sismo struct {
	ID          int
	FECHA_UTC   int
	HORA_UTC    int
	LATITUD     float64
	LONGITUD    float64
	PROFUNDIDAD int
	MAGNITUD    float64
}

func processRecord(record []string) Sismo {
	// Convertir y preprocesar los datos del registro
	id, _ := strconv.Atoi(record[0])
	fechaUTC, _ := strconv.Atoi(record[1])
	horaUTC, _ := strconv.Atoi(record[2])
	latitud, _ := strconv.ParseFloat(record[3], 64)
	longitud, _ := strconv.ParseFloat(record[4], 64)
	profundidad, _ := strconv.Atoi(record[5])
	magnitud, _ := strconv.ParseFloat(record[6], 64)

	sismo := Sismo{
		ID:          id,
		FECHA_UTC:   fechaUTC,
		HORA_UTC:    horaUTC,
		LATITUD:     latitud,
		LONGITUD:    longitud,
		PROFUNDIDAD: profundidad,
		MAGNITUD:    magnitud,
	}

	return sismo
}

func main() {
	file, err := os.Open("Catalogo196_2021.csv")
	if err != nil {
		fmt.Println("Error al leer el archivo", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer registros", err)
		return
	}

	// Omitir Header
	if len(records) > 0 {
		records = records[1:]
	}

	var preprocessedData []Sismo
	for _, record := range records {
		s := processRecord(record)
		preprocessedData = append(preprocessedData, s)
	}

	// Dividimos el 80% para el entrenamiento y el 20% para pruebas
	// splitIndex := int(0.8 * float64(len(preprocessedData)))
	// trainingData := preprocessedData[:splitIndex]
	// testingData := preprocessedData[splitIndex:]
    type NeuralNetwork struct {
        inputs       int
        hiddens      int
        outputs      int
        hiddenLayer  []float64
        outputLayer  []float64
        hiddenBias   float64
        outputBias   float64
        hiddenWeight [][]float64
        outputWeight []float64
    }
    
    func (nn *NeuralNetwork) Initialize(inputs, hiddens, outputs int) {
        nn.inputs = inputs
        nn.hiddens = hiddens
        nn.outputs = outputs
    
        nn.hiddenLayer = make([]float64, hiddens)
        nn.outputLayer = make([]float64, outputs)
    
        nn.hiddenWeight = make([][]float64, inputs)
        for i := range nn.hiddenWeight {
            nn.hiddenWeight[i] = make([]float64, hiddens)
            for j := range nn.hiddenWeight[i] {
                nn.hiddenWeight[i][j] = rand.Float64()*2 - 1 // Random value between -1 and 1
            }
        }
    
        nn.outputWeight = make([]float64, hiddens)
        for i := range nn.outputWeight {
            nn.outputWeight[i] = rand.Float64()*2 - 1
        }
    
        nn.hiddenBias = rand.Float64()*2 - 1
        nn.outputBias = rand.Float64()*2 - 1
    }
    
    func sigmoid(x float64) float64 {
        return 1.0 / (1.0 + math.Exp(-x))
    }
    
    func (nn *NeuralNetwork) Forward(inputData []float64) float64 {
        for i := 0; i < nn.hiddens; i++ {
            nn.hiddenLayer[i] = 0
            for j := 0; j < nn.inputs; j++ {
                nn.hiddenLayer[i] += inputData[j] * nn.hiddenWeight[j][i]
            }
            nn.hiddenLayer[i] += nn.hiddenBias
            nn.hiddenLayer[i] = sigmoid(nn.hiddenLayer[i])
        }
    
        for i := 0; i < nn.outputs; i++ {
            nn.outputLayer[i] = 0
            for j := 0; j < nn.hiddens; j++ {
                nn.outputLayer[i] += nn.hiddenLayer[j] * nn.outputWeight[j]
            }
            nn.outputLayer[i] += nn.outputBias
            nn.outputLayer[i] = sigmoid(nn.outputLayer[i])
        }
    
        return nn.outputLayer[0] // For simplicity, we're only returning the first output
    }
    
    func main() {
        // ... [El código anterior se mantiene igual]
    
        // Initialize neural network
        rand.Seed(time.Now().UnixNano())
        nn := &NeuralNetwork{}
        nn.Initialize(3, 5, 1) // 3 inputs, 5 hidden neurons, 1 output
    
        // Example of predicting magnitude for the first sismo in the dataset
        if len(preprocessedData) > 0 {
            input := []float64{preprocessedData[0].LATITUD, preprocessedData[0].LONGITUD, float64(preprocessedData[0].PROFUNDIDAD)}
            predictedMagnitude := nn.Forward(input)
            fmt.Println("Predicted Magnitude:", predictedMagnitude)
        }
    }
}
