# sarama kafka example

## 1. Apahe Kafka install
### 1) 다운로드 링크
    https://kafka.apache.org/downloads

![install1](https://github.com/user-attachments/assets/e7ce4e69-6bdd-4145-8de3-d84f3c7e7181)

### 2) Download Kafka using wget
* **위 이미지 우클릭 후 링크 복사**

    `wget -d https://downloads.apache.org/kafka/3.8.0/kafka_2.13-3.8.0.tgz`
<br/> 
* **kafka install file**
![install2](https://github.com/user-attachments/assets/d16daaef-6643-4e68-b889-5da26ebc7163)
    
### 3) Install Kafka
 * **3.1 압축풀기**
    `tar -zxf kafka_2.13-3.8.0.tgz`
<br/> 
 * **3.2 설치 폴더로 이동**
    `cd kafka_2.13-3.8.0/`

### 4) Kafka 실행
 * **4.1 zookeeper 실행** 
    ```
    $ cd kafka_2.13-3.8.0/bin/
    $ ./zookeeper-server-start.sh -daemon ../config/zookeeper.properties
    ```

 * **4.2 브로커 서버 설정 및 실행** 
   - **server.properties파일 수정**
   ![install3](https://github.com/user-attachments/assets/5f427025-c0f9-4295-81f5-5045052656a7)

   - **브로커 서버 실행**
   ```
    $ cd kafka_2.13-3.8.0/bin/
    $ ./kafka-server-start.sh ../config/server.properties
    ```
 
 ## 2. Kafka producer and consumer example
 ### 1) basic_consumer_produmer
 * **1.1 producer**
  - `go build producer.go and execute producer`
  
 * **1.2 consumer**
  - `go build consumer.go and execute consumer`

 * **1.3 실행화면**
 ![install4](https://github.com/user-attachments/assets/4c8015b6-bf74-4308-9dc0-cf68c129a1ad)