pipeline {
  agent any

  options {
      timeout(time: 5, unit: 'MINUTES')
      // 不允许同时执行流水线, 防止同时访问共享资源等
      disableConcurrentBuilds()
      // 显示具体的构建流程时间戳
      timestamps()
  }

  environment {
    CHANGE_LOG = sh returnStdout: true, script: 'git log --pretty=format:\'%h - %an,%ar : %s\' --since=\'1 hours\' | head -n 1'
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
      post {
        success {
          sh 'docker rmi `docker images | awk \'/^<none>/ { print $3 }\'`'
        }
      }
    }
  }

  post {
    always {
      withCredentials([string(credentialsId: 'PUSH_KEY', variable: 'PUSH_KEY')]) {
        echo "text=${currentBuild.projectName} 集成 ${currentBuild.currentResult} ${env.CHANGE_LOG}"
        sh "curl -d 'text=${currentBuild.projectName} 集成 ${currentBuild.result}' -d 'desp=change log: `${env.CHANGE_LOG}`' 'https://sc.ftqq.com/${PUSH_KEY}.send'"
        // httpRequest consoleLogResponseBody: true, httpMode: 'POST',  contentType: 'APPLICATION_FORM', requestBody: "text=${currentBuild.projectName} 集成 ${currentBuild.currentResult}&desp=change: `${currentBuild.description}`", responseHandle: 'NONE', url: "https://sc.ftqq.com/${PUSH_KEY}.send"
      }
    }
  }
}