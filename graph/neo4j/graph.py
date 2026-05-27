from neo4j import GraphDatabase

class TopologyGraph:

    def __init__(self, uri, user, password):

        self.driver = GraphDatabase.driver(
            uri,
            auth=(user, password)
        )