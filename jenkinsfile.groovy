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
                sh'''
                   echo ${WORKSPACE}
                   cd  ${WORKSPACE}
                   go build 
                  '''
             }
         }
         stage('停止项目') {
             steps {
             echo '停止项目'
             }
         }
         stage('运行项目') {
             steps {
                echo '复制项目'
             }
         }
    }

}