# Tài liệu Triển khai Dự án

README này cung cấp hướng dẫn về cách thiết lập và chạy pipeline Jenkins để triển khai hình ảnh Docker `ngocthuong/server-final`.

## Điều kiện tiên quyết

Trước khi chạy pipeline, hãy đảm bảo bạn đã có:

- **Jenkins** được cài đặt và cấu hình. 
- **Docker** được cài đặt trên máy chủ Jenkins. 
- **Git** được cài đặt trên máy chủ Jenkins. 
- Tài khoản **Docker Hub** và thông tin đăng nhập được lưu trữ trong Jenkins dưới dạng `docker-hub-credentials`.
- Một **EC2 instance AWS** được thiết lập để triển khai sản xuất, với quyền truy cập SSH được cấu hình.
- Một **Bot Telegram** đã được tạo với token và chat ID được thiết lập trong các biến môi trường.
- [Hướng dẫn triển khai Jenkins với Docker trên AWS EC2](https://viblo.asia/p/cach-trien-khai-mot-du-an-bang-jenkins-docker-ec2-3kY4gnD0VAe)
- [Tạo CI/CD Pipeline với Jenkins trên Amazon ECS](https://locker.io/vi/blog/cach-tao-ci-cd-pipeline-voi-jenkins)
- [Hướng dẫn cài đặt Jenkins trên AWS](https://www.jenkins.io/doc/tutorials/tutorial-for-installing-jenkins-on-AWS/)

## Hướng dẫn cài đặt Jenkins với Docker
  ```
$ docker pull jenkins/jenkins:lts
$ docker run -d --name jenkins \
    -p 8080:8080 -p 50000:50000 \
    -v jenkins_home:/var/jenkins_home \
    -v /var/run/docker.sock:/var/run/docker.sock \
    jenkins/jenkins:lts
$ docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```
```
$ docker exec -it --user root jenkins /bin/bash
$ apt-get update
$ apt-get install -y docker.io
$ apt-get install -y docker-ce docker-ce-cli containerd.io
$ docker --version
$ usermod -aG docker jenkins
$ docker restart jenkins
$ chmod 666 /var/run/docker.sock
```
## Biến môi trường

Hãy chắc chắn thiết lập các biến môi trường sau trong pipeline Jenkins của bạn:

- `DOCKER_IMAGE`: Tên hình ảnh Docker (ví dụ: `ngocthuong/server-final`)
- `DOCKER_TAG`: Nhãn hình ảnh Docker (mặc định là `latest`)
- `TELEGRAM_BOT_TOKEN`: Token của Bot Telegram của bạn
- `TELEGRAM_CHAT_ID`: ID chat để gửi thông báo
- `PROD_SERVER`: DNS công khai hoặc IP của EC2 instance AWS của bạn
- `PROD_USER`: Tên người dùng để truy cập SSH (mặc định là `ubuntu`)

## Các giai đoạn của Pipeline

Pipeline Jenkins bao gồm các giai đoạn sau:

1. **Clone Repository**: Sao chép kho lưu trữ từ GitHub.
2. **Build Docker Image**: Xây dựng hình ảnh Docker bằng cách sử dụng Dockerfile trong kho lưu trữ.
3. **Run Tests**: Nơi dành cho việc chạy các bài kiểm tra (tùy chỉnh theo nhu cầu).
4. **Push to Docker Hub**: Đẩy hình ảnh Docker đã xây dựng lên Docker Hub.
5. **Deploy Golang to DEV**: Triển khai ứng dụng lên môi trường phát triển.
6. **Deploy to Production on AWS**: Triển khai ứng dụng lên môi trường sản xuất trên AWS EC2.

## Chạy Pipeline

1. **Thiết lập Job Jenkins**:
   - Tạo một job pipeline mới trong Jenkins.
   - Sao chép và dán Jenkinsfile đã cung cấp vào cấu hình job.

2. **Kích hoạt Pipeline**:
   - Kích hoạt thủ công job hoặc cấu hình để nó chạy khi có thay đổi SCM.

## Thông báo

Pipeline sẽ gửi thông báo tới một chat Telegram được chỉ định khi build thành công hoặc thất bại. Hãy đảm bảo rằng hàm `sendTelegramMessage` được cấu hình chính xác với token Bot Telegram và chat ID của bạn.

## Dọn dẹp

Sau khi pipeline hoàn thành, nó sẽ tự động dọn dẹp không gian làm việc để đảm bảo không còn tệp tin dư thừa nào còn lại.

## Khó khăn

- Đảm bảo rằng tất cả các biến môi trường đã được thiết lập chính xác.
- Xác minh rằng instance Jenkins có quyền truy cập vào Docker và có thể giao tiếp với EC2 instance AWS.
- Kiểm tra đầu ra của console Jenkins để tìm bất kỳ lỗi nào trong quá trình thực thi pipeline.

## Kết luận

Pipeline này tự động hóa quá trình xây dựng, kiểm tra và triển khai ứng dụng. Tùy chỉnh giai đoạn kiểm tra và bất kỳ cấu hình nào khác theo nhu cầu cụ thể của dự án của bạn.
