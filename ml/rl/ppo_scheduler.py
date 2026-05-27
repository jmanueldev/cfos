import torch
import torch.nn as nn

class Policy(nn.Module):

    def __init__(self, state, actions):
        super().__init__()

        self.net = nn.Sequential(
            nn.Linear(state, 512),
            nn.ReLU(),
            nn.Linear(512, 256),
            nn.ReLU(),
            nn.Linear(256, actions)
        )

    def forward(self, x):
        return self.net(x)