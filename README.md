# Go FRED Scraper

`go-fred-scraper`, [FRED (Federal Reserve Economic Data)](https://fred.stlouisfed.org/) API'sini kullanarak belirli bir ekonomik serinin verilerini istediğiniz tarih aralığında çekip CSV formatında bir dosyaya kaydetmenizi sağlayan bir komut satırı aracıdır (CLI).

Bu araç, Go programlama dili ile SOLID prensiplerine uygun ve modüler bir yapıda geliştirilmiştir.

## Özellikler

- Belirtilen FRED serisi için veri çekme
- Başlangıç ve bitiş tarihlerine göre veri aralığını filtreleme
- Çekilen verileri `.csv` uzantılı bir dosyaya kaydetme
- Kolay ve anlaşılır komut satırı arayüzü

## Kurulum

Projeyi kullanmaya başlamadan önce Go'nun sisteminizde kurulu olduğundan emin olun.

1.  **Projeyi Klonlayın (İsteğe Bağlı):**

    ```bash
    git clone https://github.com/mr-isik/go-fred-scraper.git
    cd go-fred-scraper
    ```

2.  **Bağımlılıkları Yükleyin:**

    ```bash
    go mod tidy
    ```

3.  **`.env` Dosyasını Oluşturun:**

    Uygulamanın FRED API'si ile iletişim kurabilmesi için bir API anahtarına ihtiyacı vardır. FRED web sitesinden ücretsiz olarak bir API anahtarı alabilirsiniz.

    Projenin ana dizininde `.env` adında bir dosya oluşturun ve içine API anahtarınızı aşağıdaki gibi ekleyin:

    ```
    FRED_API_KEY=SENIN_API_ANAHTARIN
    ```

    `SENIN_API_ANAHTARIN` kısmını kendi FRED API anahtarınız ile değiştirmeyi unutmayın.

## Kullanım

Uygulamayı çalıştırmak için `go run` komutunu kullanabilir veya projeyi derleyerek çalıştırılabilir bir dosya oluşturabilirsiniz.

### Parametreler

- `-series`: (Zorunlu) Verisini çekmek istediğiniz FRED serisinin ID'si. (Örn: `GNPCA`, `DEXUSEU`)
- `-start`: (Zorunlu) Veri çekilecek başlangıç tarihi. `YYYY-MM-DD` formatında olmalıdır.
- `-end`: (Zorunlu) Veri çekilecek bitiş tarihi. `YYYY-MM-DD` formatında olmalıdır.
- `-output`: (İsteğe bağlı) Oluşturulacak CSV dosyasının adı. Varsayılan değer `output.csv`'dir.

### `go run` ile Çalıştırma

Proje dizinindeyken aşağıdaki komutu kullanarak uygulamayı çalıştırabilirsiniz:

```bash
go run ./cmd/go-fred-scraper -series GNPCA -start 2000-01-01 -end 2023-01-01 -output gnp_data.csv
```

Bu komut, `GNPCA` serisinin 1 Ocak 2000 ile 1 Ocak 2023 arasındaki verilerini çeker ve proje ana dizininde `gnp_data.csv` adında bir dosyaya kaydeder.

### Projeyi Derleme (Build)

Uygulamayı her seferinde `go run` ile çalıştırmak yerine, derleyerek tek bir çalıştırılabilir dosya haline getirebilirsiniz.

```bash
go build -o fred-scraper ./cmd/go-fred-scraper
```

Bu komut, proje ana dizininde `fred-scraper` (Windows için `fred-scraper.exe`) adında bir çalıştırılabilir dosya oluşturur.

### Derlenmiş Uygulamayı Çalıştırma

**Windows (PowerShell):**

```powershell
./fred-scraper.exe -series DEXUSEU -start 2022-01-01 -end 2023-01-01 -output dolar_euro_kuru.csv
```

**Linux / macOS:**

```bash
./fred-scraper -series DEXUSEU -start 2022-01-01 -end 2023-01-01 -output dolar_euro_kuru.csv
```

Bu komut, Dolar/Euro paritesinin belirtilen tarih aralığındaki verilerini çeker ve `dolar_euro_kuru.csv` dosyasına yazar.

## Proje Yapısı

Proje, sorumlulukların ayrılması (Separation of Concerns) ilkesine uygun olarak modüler bir yapıda tasarlanmıştır.

```
go-fred-scraper/
├── cmd/go-fred-scraper/
│   └── main.go         # CLI komutları, flag yönetimi ve ana uygulama mantığı
├── internal/
│   ├── client/
│   │   └── client.go   # FRED API istemcisini oluşturan ve yöneten kod
│   ├── series/
│   │   └── series.go   # FRED serileriyle ilgili iş mantığı
│   └── writer/
│       └── writer.go   # Verileri CSV formatında dosyaya yazan kod
├── pkg/
│   ├── config/
│   │   └── config.go   # Uygulama yapılandırması (örn. API anahtarı)
│   └── fred/
│       └── fred.go     # FRED API ile doğrudan iletişim kuran istemci
├── go.mod
├── go.sum
└── README.md
```

## Lisans

Bu proje MIT Lisansı altında lisanslanmıştır.
