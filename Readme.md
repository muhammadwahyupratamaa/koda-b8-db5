# ERD Data Contact List

## Berikut merupakan tampilan ERD dengan mermaid untuk contact list

### Akses Aplikasi
Untuk menjalankan aplikasi bisa pull dari package dan coba jalankan dengan -it supaya aplikasi berjalan di terminal

```sh
docker pull ghcr.io/muhammadwahyupratamaa/contact-list:latest
docker run --rm -it ghcr.io/muhammadwahyupratamaa/contact-list:latest
```

```mermaid

erDiagram

CONTACTS{
    int id PK
    string nama
    string email
    string phone
    timestamp created_At
    timestamp updated_at
}

```

