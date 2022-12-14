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
                sh  '''
                   echo ${PATH}
                   echo ${WORKSPACE}
                   cd  ${WORKSPACE}
                   go build -o /home/study/memoirs
                  '''
            }
        }
        stage('停止项目') {
            steps {
                sh ''' 
                    pid=$(netstat -antp | grep :8888 | awk '{print $7}' | awk -F'/' '{ print $1 }')
                    echo 当前进程:$pid
                    if [ -n "${pid}" ] ; then
                        echo "kill -9 pid:${pid}"
                        kill -9 $pid
                    fi
                    sleep 5
                   '''
            }
        }
        stage('运行项目') {
            steps {
                withEnv(['JENKINS_NODE_COOKIE=dontkillme']) {
                    sh '''               
                    nohup /home/study/memoirs >/home/study/memoirs.log 2>&1 &
                  '''
                }
            }
        }
    }

}