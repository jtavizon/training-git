amigoscode.com
   Git and Github Essentials
   1) Download and install github https://github.com/git-guides/install-git
   2) cmd
   3) git config --global user.name "jtavizon"
      git config --global user.email "jairo.tavizon@gmail.com"
      git config --global color.ui auto
      git config -l

   4) Initialising git repository
      NEW projects
          cd C:\ProgramFilesTavi\Training
          git init .

   5) git add
       Create files 
       git add githubLearning                    to add into stage area
       git rm --cached githubLearning            to remove from stage area
       git add .                                 to add everything to stage area
       git status

   6) git commit 
       git commit -m "proyecto aprender github"
       git log
       git show 0292215647839282eeba506bab4a62a3348c68a9
       git status
       git diff                                 para ver diferencias 
       git commit --amend -m "para ver diferencias mas sentido"

   7) push
       git remote add origin https://github.com/jtavizon/training-git.git
       git branch -M main
       click logo, settings, SSH keys GPG keys, generating SSH keys instructions
              https://docs.github.com/en/authentication/connecting-to-github-with-ssh
              windows
        en git command
        ssh-keygen -t ed25519 -C "jairo.tavizon@gmail.com"
        eval "$(ssh-agent -s)"
        ssh-add ~/.ssh/id_ed25519
        clip SSH         ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIMMFwUM4of1z4RCpbijAB+MkgOMA0oHi/t6zZvKepcls jairo.tavizon@gmail.com
        git push -u origin main

   8) Ya se debe ver en github.com









       