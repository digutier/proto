# proto

Aqui va el proto
message Mensaje {
    string body = 1;
}

message OrdenPyme {
    string id = 1;
    string producto = 2;
    int64 valor = 3;
    string tienda = 4;
    string destino = 5;
    int64 prioridad = 6;
}

message Codigo {
    int64 cod = 1; 
}
// Definir el paquete
message Paquete {
    
}
service CourierService {
    rpc Hello(Mensaje) returns (Mensaje) {}
    rpc CodigoPyme(OrdenPyme) returns (Codigo) {}
}
