import torch
import torch.nn as nn

class DemandPredictor(nn.Module):

    def __init__(self):
        super().__init__()

        self.lstm = nn.LSTM(
            input_size=32,
            hidden_size=128,
            num_layers=3,
            batch_first=True
        )

        self.fc = nn.Linear(128, 1)

    def forward(self, x):

        out, _ = self.lstm(x)

        return self.fc(out[:, -1, :])