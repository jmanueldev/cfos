import grpc
import cfos_pb2
import cfos_pb2_grpc

channel = grpc.insecure_channel("localhost:50051")

client = cfos_pb2_grpc.CognitiveFabricStub(channel)

response = client.SubmitIntent(
    cfos_pb2.IntentRequest(
        workload_id="llm-train-001",
        objective="minimize_carbon",
    )
)

print(response.execution_plan_id)