package main
import (
	"fmt"
)

const (
  PassStatus = "pass"
  FailStatus = "fail"
 )

type HealthCheck struct{
	ServiceID int
	Status string
}

func GenerateCheck() []HealthCheck {
	services := []HealthCheck{}
	for i := 0; i < 5; i ++ {
		if i % 2 == 0 {
			services = append(services, HealthCheck{
				ServiceID: i,
				Status: PassStatus,
			})
		} else {
			services = append(services, HealthCheck{
				ServiceID: i,
				Status: FailStatus,
			})
		}
	}
	return services
}

func main() {
	fmt.Println("Тут будет выведен идентификатор")
	services := GenerateCheck()

	// Выводим результат
	for _, service := range services {
		if service.Status == PassStatus {
			fmt.Println(service.ServiceID)	
		}
		
	}
}

