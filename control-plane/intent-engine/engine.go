package intent

type Intent struct {
    WorkloadID string
    Objective  string
    Constraints map[string]string
}

type Engine struct {}

func New() *Engine {
    return &Engine{}
}

func (e *Engine) Parse(text string) (*Intent, error) {

    return &Intent{
        WorkloadID: "wf-001",
        Objective: "minimize_latency",
        Constraints: map[string]string{
            "region": "us-east",
        },
    }, nil
}