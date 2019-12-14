pipeline {
  agent any

  options {
      timeout(time: 5, unit: 'MINUTES')
      // 不允许同时执行流水线, 防止同时访问共享资源等
      disableConcurrentBuilds()
      // 显示具体的构建流程时间戳
      timestamps()
  }

  stages {
    stage('Test') {
      agent {
        docker {
          image 'golang:1.13'
        }
      }
      steps {
        sh 'go test -v ./test'
      }
    }

    stage('Build Docker') {
      steps {
        sh 'docker build -t advproject .'
      }
    }
  }

  post {
    always {
      withCredentials([string(credentialsId: 'PUSH_KEY', variable: 'PUSH_KEY')]) {
        sh '''
          curl -d "text=${currentBuild.projectName} 集成 ${currentBuild.result}" -d "desp=change: `${CHANGE_TITLE}`" "https://sc.ftqq.com/${PUSH_KEY}.send"
        '''
      }
    }
  }
}