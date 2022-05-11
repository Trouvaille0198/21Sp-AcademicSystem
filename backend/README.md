# 后端
## TODO
- 老师
    - 查看各学期其教授的开课课程
    - 查看课程里的学生
    - 登记、修改成绩（平时分+考试分）
- 学生
    - 查看已选课程
        - 退课
    - 查看成绩
    - 查看所有开课课程
        - 选课
- 管理员
    - 查看学生
        - 增删学生
    - 查看老师
        - 增删老师
    - 查看课程
        - 增删课程

## Docker 使用方法
安装 Docker
```shell
sudo apt-get install docker-ce docker-ce-cli containerd.io
```
构建 image
```shell
docker build -t go_app .
```

运行 Docker
```shell
docker run -p 8080:8080 go_app
```

