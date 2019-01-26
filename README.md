# Productivity Challenge Timer Backup Reader

## Productivity Challenge... what?

This is a simple script to extract backup data from the app [Productivity Challenge Timer](https://play.google.com/store/apps/details?id=com.wlxd.pomochallenge&hl=en) on Android.

It's basically an app which give you ranks depending on how much Pomodoro you're doing. It's a great app if you need something which push you gently to do Pomodoro and therefore accomplish your goals.

## What does it extract?

Productivity Challenge Timer Backup Reader will group every Pomodoro you did per week and per project. When a project has `\\` in its name, it's a separator between the parent project and the child project.

## How does it work?

1. Download the binary. From a terminal using curl:

`cd ~/Documents && sudo curl -LO https://github.com/Phantas0s/productivity-challenge-reader/releases/download/0.1.0/productivity-challenge-reader && cd -`

If you want to install it somewhere else than `~/Documents/`, just change this part of the command (after the first `cd` and before `&&`)

You can as well download the source and compile it by yourself. It will take less than 2 seconds.

2. Create a backup from the Productivity Challenge Timer app
3. Copy the backup to wherever you installed Productivity Challenge Timer Backup Reader (in my example `~/Documents/`)
4. Rename the backup file you copied to `backup.json`
5. Run the Productivity Challenge Timer Backup Reader from the terminal: `./productivity-challenge-reader`
6. The output will look like that:

```
2018-10-15,Mathematics \\ Arithmetic,2
2018-10-22,Mathematics \\ Arithmetic,2
2018-10-29,Mathematics \\ Arithmetic,2
2018-11-12,Mathematics \\ Arithmetic,1
2018-09-17,Mathematics \\ Arithmetic,10
```

It means that during the week beginning the Monday `2018-10-15`, I did 2 Pomodoros for the project `Arithmetic`, which has for parent `Mathematics`.

7. If you want to visualise it with LibreOffice or (fool!) Microsoft Excel, you can do that:
`./productivity-challenge-reader > feelmypomodoropower.csv`

Then open `feelmypomodoropower.csv` with whatever sheet program you want.

## Contributing

You know. Pull requests.

## Licence

Do whatever you want with that :D 


