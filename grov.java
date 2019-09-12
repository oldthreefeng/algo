
node {
    parallel 'check one': {
         def remote = [:]
         remote.name = "qx-file"
         remote.host = "10.9.152.44"
         remote.port = 6822
         remote.allowAnyHosts = true
         withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
                remote.user = username
                remote.identityFile = identity
                remote.passphrase = passphrase
                stage("SSH Steps Rocks!") {
                    sshCommand remote: remote, command: 'sh /data/qianxiang/build.sh'
                }
         }
    }, 'check two': {
         def remote = [:]
         remote.name = "qx-file"
         remote.host = "10.9.38.194"
         remote.port = 6822
         remote.allowAnyHosts = true
         withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
                remote.user = username
                remote.identityFile = identity
                remote.passphrase = passphrase
                stage("SSH Steps Rocks!") {
                    sshCommand remote: remote, command: 'sh /data/qianxiang/build.sh'
                }
         }
    }

    'check three': {
         echo "finish three"
    }
}



def remote = [:]
remote.name = "qx-file"
remote.host = "10.9.152.44"
remote.port = 6822
remote.allowAnyHosts = true

node {

    withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
        remote.user = username
        remote.identityFile = identity
        remote.passphrase = passphrase
        stage("SSH Steps Rocks!") {
            sshCommand remote: remote, command: 'for i in {1..5}; do echo -n \"Loop \$i \"; date ; sleep 1; done'
        }
    }
}


node {
    stage('Parallel Stage') {
        parallel 'qianxiang-file': {
                def remote = [:]
                remote.name = "qx-file"
                remote.host = "10.9.152.44"
                remote.port = 6822
                remote.allowAnyHosts = true
                withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
                remote.user = username
                remote.identityFile = identity
                remote.passphrase = passphrase
                stage("SSH Steps Rocks!") {
                sshCommand remote: remote, command: 'sh /data/qianxiang/build.sh'
                }
            }
        }, 'qianxiang-cron': {
                def remote = [:]
                remote.name = "qx-cron"
                remote.host = "10.9.38.194"
                remote.port = 6822
                remote.allowAnyHosts = true
                withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
                remote.user = username
                remote.identityFile = identity
                remote.passphrase = passphrase
                stage("SSH Steps Rocks!") {
                sshCommand remote: remote, command: 'sh /data/qianxiang/build.sh'
                }
            }
        }
    }
    stage('DingDing') {
         steps {
             echo 'This stage will be executed after success.'
         }
    }
}

pipeline {
    agent any
    stages {
        stage('Parallel Stage') {
            parallel {
                stage('Stage1.1') {
                    def remote = [:]
                    remote.name = "qx-file"
                    remote.host = "10.9.152.44"
                    remote.port = 6822
                    remote.allowAnyHosts = true
                    withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
                        remote.user = username
                        remote.identityFile = identity
                        remote.passphrase = passphrase
                        stage("qx-file-build") {
                        sshCommand remote: remote, command: 'uname -a'
                        }
                    }
                }
                stage ('1.2'){
                    def remote = [:]
                    remote.name = "qx-cron"
                    remote.host = "10.9.38.194"
                    remote.port = 6822
                    remote.allowAnyHosts = true
                    withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
                        remote.user = username
                        remote.identityFile = identity
                        remote.passphrase = passphrase
                        stage("qx-cron-build") {
                        sshCommand remote: remote, command: 'uname -a'
                        }
                    }
                }
            }

        }
    }
    post {
        success {
            dingTalk accessToken: 'c580eda8baa9e1df638c20043f1009b25dd12a9554590a662571bd7ed2f14f07',
            imageUrl: '',
            jenkinsUrl: '',
            message: '构建成功 ',
            notifyPeople: ''
        }
        failure {
            dingTalk accessToken: 'c580eda8baa9e1df638c20043f1009b25dd12a9554590a662571bd7ed2f14f07',
            imageUrl: '',
            jenkinsUrl: '',
            message: '构建失败',
            notifyPeople: ''
        }

    }
}



node {
    stage('Parallel Stage') {
        parallel 'qianxiang-file': {
                def remote = [:]
                remote.name = "qx-file"
                remote.host = "10.9.152.44"
                remote.port = 6822
                remote.allowAnyHosts = true
                withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
                remote.user = username
                remote.identityFile = identity
                remote.passphrase = passphrase
                stage("qx-file-build") {
                sshCommand remote: remote, command: 'uname -a'
                }
            }
        }, 'qianxiang-cron': {
                def remote = [:]
                remote.name = "qx-cron"
                remote.host = "10.9.38.194"
                remote.port = 6822
                remote.allowAnyHosts = true
                withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
                remote.user = username
                remote.identityFile = identity
                remote.passphrase = passphrase
                stage("qx-cron-build") {
                sshCommand remote: remote, command: 'uname -a'
                }
            }
        }
    }
}

    post {
        success {
            dingTalk accessToken: 'c580eda8baa9e1df638c20043f1009b25dd12a9554590a662571bd7ed2f14f07',
            imageUrl: '',
            jenkinsUrl: '',
            message: '构建成功 ',
            notifyPeople: ''
        }
        failure {
            dingTalk accessToken: 'c580eda8baa9e1df638c20043f1009b25dd12a9554590a662571bd7ed2f14f07',
            imageUrl: '',
            jenkinsUrl: '',
            message: '构建失败',
            notifyPeople: ''
        }

    }
pipeline {
  agent any
    stages {
      stage ('Main Stage') {
        steps {
          script {
            stage ('nothing but display a message') {
              echo 'First stage'
}
            stage ('ssh') {
              withCredentials([sshUserPrivateKey(credentialsId: 'xxxxxxxxxxxxx',   keyFileVariable: 'key', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
{{ remote.passphrase = passphrase}}
                sshagent(['xxxxxxxxxxxxx']) {
   sh 'scp -o StrictHostKeyChecking=no $username@hostname:~/abc.sh .'
}
}
}
}
}
}
}
}


pipeline {
    agent any
    stages {
        stage('并行执行的 Stage') {
            parallel {
                stage('Main Stage') {
                    steps {
                        script {
                            stage ('ssh') {
                                withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
                                {{ remote.passphrase = passphrase}}
                                    sshagent(['louis']) {
                                    ssh -t -p 6822 root@10.9.38.194 "uname -r"
                                    }
                                }
                            }
                        }
                    }
                }
                stage('Stage2.2') {
                    steps {
                        script {
                            stage ('ssh') {
                                withCredentials([sshUserPrivateKey(credentialsId: 'louis', keyFileVariable: 'identity', passphraseVariable: 'passphrase', usernameVariable: 'username')]) {
                                    {{remote.passphrase = passphrase}}
                                    sshagent(['louis']) {
                                    ssh -t -p 6822 root@10.9.152.44 "uname -r"
                                    }
                                }
                            }
                        }

                    }
                }
            }
        }
    }
    post {
        success {
            dingTalk accessToken: 'c580eda8baa9e1df638c20043f1009b25dd12a9554590a662571bd7ed2f14f07',
            imageUrl: '',
            jenkinsUrl: '',
            message: '构建成功 ',
            notifyPeople: ''
        }
        failure {
            dingTalk accessToken: 'c580eda8baa9e1df638c20043f1009b25dd12a9554590a662571bd7ed2f14f07',
            imageUrl: '',
            jenkinsUrl: '',
            message: '构建失败',
            notifyPeople: ''
        }

    }
}