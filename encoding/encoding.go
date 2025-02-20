package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	jsonData, err := os.ReadFile(j.FileInput)
	if err != nil {
		return fmt.Errorf("ошибка при  чтении файла при преобразовании JSON в YAML %w", err)
	}

	var dc models.DockerCompose

	if err = json.Unmarshal(jsonData, &dc); err != nil {
		return fmt.Errorf("ошибка при анмаршелинге json в JSON в YAML %w", err)
	}

	yamlData, err := yaml.Marshal(dc)
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании JSON в YAML %w", err)
	}

	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла ymal при JSON в YAML %w", err)
	}

	defer yamlFile.Close()

	_, err = yamlFile.Write(yamlData) 
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл yamlFile при JSON в YAML %w", err)
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlData, err := os.ReadFile(y.FileInput)
	if err != nil {
		return fmt.Errorf("ошибка при  чтении файла при преобразовании YAML в JSON %w", err)
	}

	var dc models.DockerCompose

	if err = yaml.Unmarshal(yamlData, &dc); err != nil {
		return fmt.Errorf("ошибка при анмаршелинге yaml в YAML в JSON %w", err)
	}

	jsonData, err := json.Marshal(dc)
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании YAML в JSON %w", err)
	}

	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла ymal при YAML в JSON %w", err)
	}

	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData) 
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл yamlFile при YAML в JSON %w", err)
	}

	return nil
}
