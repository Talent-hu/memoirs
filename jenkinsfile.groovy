pipeline {
    agent any
    stages {

        stage('拉取文件') {
            steps {
                git credentialsId: '18387188346', url: 'https://gitee.com/hutiancai/memo.git'
            }
        }
        stage('编译代码') {
            steps {
                sh '''
                   echo ${WORKSPACE}
                   cd  ${WORKSPACE}
                   go build -o /home/study/memoirs
                  '''
            }
        }
        stage('停止项目') {
            steps {
                sh ''' 
                    kill -9 $(netstat -antp | grep :8888 | awk '{print $7}' | awk -F'/' '{ print $1 }') 
                   '''
            }
        }
        stage('运行项目') {
            steps {
                sh ''' nohup /home/study/memoirs > memoirs.log & '''
            }
        }
    }

}