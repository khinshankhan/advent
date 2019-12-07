import networkx as nx

def graphy(f, val1, val2):
    g = nx.Graph(i.strip().split(")") for i in f)

    val1 = sum(nx.shortest_path_length(g, "COM").values())
    val2 = nx.shortest_path_length(g, "YOU", "SAN") - 2

    return (val1, val2)
if __name__ == "__main__":
    f = open("input.txt")
    print(graphy(f, 0, 0))
    f.close()
