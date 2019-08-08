(setq inhibit-startup-message t) ;close start page
(desktop-save-mode t);保存session

(defun plist-to-alist (the-plist)
(defun get-tuple-from-plist (the-plist)
(when the-plist (cons (car the-plist) (cadr the-plist))))
(let ((alist *()))
(while the-plist
(add-to-list *alist (get-tuple-from-plist the-plist))
(setq the-plist (cddr the-plist))) alist))

(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(global-linum-mode t)
)
(custom-set-faces
 ;; custom-set-faces was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
)
(add-to-list 'load-path "~/.emacs.d/elpa/color-theme-20080305.34/")
(require 'color-theme)
(color-theme-initialize)
(color-theme-pok-wog)

(require 'package) ;; You might already have this line
(add-to-list 'package-archives
             '("melpa" . "https://melpa.org/packages/"))
(when (< emacs-major-version 24)
  ;; For important compatibility libraries like cl-lib
  (add-to-list 'package-archives '("gnu" . "http://elpa.gnu.org/packages/")))
(package-initialize) ;; You might already have this line

;evil
(require 'evil)
(evil-mode 1)

;auto-complete
(require 'auto-complete-config)
(ac-config-default)

					;neotree
(require 'neotree)
(global-set-key (kbd "<f5>") 'neotree-toggle)
(global-set-key (kbd "<f6>")
		(lambda ()
		  (interactive)
		  (neotree-dir "/var/www/html")
		  ))
(defun neotree-close-up-node ()
  (interactive)
  (neotree-select-up-node)
  (neotree-enter))

(defun neotree-current-to-root ()
  (interactive)
  (neotree-change-root))

(defun neotree-open-vsplite ()
  (interactive)
  (neotree-enter-vertical-split))

(defun neotree-open-splite ()
  (interactive)
  (neotree-enter-horizontal-split))

(setq neo-smart-open t)
(add-hook 'neotree-mode-hook
	(lambda ()
	    (define-key evil-normal-state-local-map (kbd "TAB") 'neotree-enter)
	    (define-key evil-normal-state-local-map (kbd "o") 'neotree-enter)
	    (define-key evil-normal-state-local-map (kbd "i") 'neotree-open-splite)
	    (define-key evil-normal-state-local-map (kbd "s") 'neotree-open-vsplite)
	    (define-key evil-normal-state-local-map (kbd "t") 'neotree-current-to-root)
	    (define-key evil-normal-state-local-map (kbd "x") 'neotree-close-up-node)
	    (define-key evil-normal-state-local-map (kbd "SPC") 'neotree-enter)
	    (define-key evil-normal-state-local-map (kbd "q") 'neotree-hide)
	    (define-key evil-normal-state-local-map (kbd "RET") 'neotree-enter)))
(electric-pair-mode t)
(require 'highlight-parentheses)  
(define-globalized-minor-mode global-highlight-parentheses-mode  
  highlight-parentheses-mode  
  (lambda ()  
    (highlight-parentheses-mode t)))  
(global-highlight-parentheses-mode t) 

(require 'php-mode)
(add-to-list 'auto-mode-alist '("\\.php[34]?\\'\\|\\.phtml\\'" . php-mode))
(add-to-list 'auto-mode-alist '("\\.module\\'" . php-mode))
(add-to-list 'auto-mode-alist '("\\.inc\\'" . php-mode))
(define-key php-mode-map                                                                            
  [(tab)]                                                                                   
  'php-complete-function)

					;(setq auto-complete-mode t)
					;(setq ac-auto-start 3) ;; 输入4个字符才开始补全
					;(global-set-key "\M-/" 'auto-complete)
					;(global-set-key [tab] 'auto-complete)
					;(setq default-major-mode 'text-mode)

(setq column-number-mode t)
(setq line-number-mode t)

(require 'flymake)

(defun flymake-php-init ()
  "Use php to check the syntax of the current file."
  (let* ((temp (flymake-init-create-temp-buffer-copy 'flymake-create-temp-inplace))
	 (local (file-relative-name temp (file-name-directory buffer-file-name))))
    (list "php" (list "-f" local "-l"))))

(add-to-list 'flymake-err-line-patterns 
  '("\\(Parse\\|Fatal\\) error: +\\(.*?\\) in \\(.*?\\) on line \\([0-9]+\\)$" 3 4 nil 2))

(add-to-list 'flymake-allowed-file-name-masks '("\\.php$" flymake-php-init))

;; Drupal-type extensions
(add-to-list 'flymake-allowed-file-name-masks '("\\.module$" flymake-php-init))
(add-to-list 'flymake-allowed-file-name-masks '("\\.install$" flymake-php-init))
(add-to-list 'flymake-allowed-file-name-masks '("\\.inc$" flymake-php-init))
(add-to-list 'flymake-allowed-file-name-masks '("\\.engine$" flymake-php-init))

(add-hook 'php-mode-hook (lambda () (flymake-mode 1)))
(define-key php-mode-map '[M-S-up] 'flymake-goto-prev-error)
(define-key php-mode-map '[M-S-down] 'flymake-goto-next-error)
