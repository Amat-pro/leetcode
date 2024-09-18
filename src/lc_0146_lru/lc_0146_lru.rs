use std::{cell::RefCell, collections::HashMap, rc::Rc};

pub struct LRUCache {
    capacity: i32,
    len: i32,
    // 哨兵 - sentinel
    head: Rc<RefCell<Node>>,
    tail: Rc<RefCell<Node>>,
    key_to_node: HashMap<i32, Rc<RefCell<Node>>>,
}

impl LRUCache {
    fn new(capacity: i32) -> Self {
        let head = Node::new(0, 0);
        let tail = Node::new(0, 0);
        head.borrow_mut().next = Some(tail.clone());
        tail.borrow_mut().pre = Some(head.clone());

        let cache = LRUCache {
            capacity,
            len: 0,
            head,
            tail,
            key_to_node: HashMap::with_capacity(capacity as usize),
        };

        return cache;
    }

    fn get(&self, key: i32) -> i32 {
        if let Some(node) = self.key_to_node.get(&key) {
            self.move_at(node.clone(), self.head.clone());
            return node.borrow().value;
        }

        -1
    }

    fn put(&mut self, key: i32, val: i32) {
        if let Some(node) = self.key_to_node.get(&key) {
            node.borrow_mut().value = val;
            self.move_at(node.clone(), self.head.clone());
            return;
        }

        if self.len > 1 && self.len == self.capacity {
            self.remove_tail();
        }

        self.push_front(Node::new(key, val));
        self.len += 1;
    }

    fn remove_tail(&mut self) {
        let tail = self.tail.clone();
        let to_remove = tail.borrow().pre.clone().unwrap();
        let pre = to_remove.borrow().pre.clone().unwrap();

        tail.borrow_mut().pre = Some(pre.clone());
        pre.borrow_mut().next = Some(self.tail.clone());

        to_remove.borrow_mut().next = None;
        to_remove.borrow_mut().pre = None;

        self.key_to_node.remove(&to_remove.borrow().key);

        self.len -= 1;
    }

    fn move_at(&self, node: Rc<RefCell<Node>>, at: Rc<RefCell<Node>>) {
        if Rc::ptr_eq(&node, &at) {
            return;
        }

        let node_next = node.borrow().next.clone().unwrap();
        let node_pre = node.borrow().pre.clone().unwrap();

        node_next.borrow_mut().pre = Some(node_pre.clone());
        node_pre.borrow_mut().next = Some(node_next.clone());

        let at_next = at.borrow().next.clone().unwrap();
        at.borrow_mut().next = Some(node.clone());
        node.borrow_mut().pre = Some(at.clone());
        node.borrow_mut().next = Some(at_next.clone());
        at_next.borrow_mut().pre = Some(node.clone());
    }

    fn push_front(&mut self, node: Rc<RefCell<Node>>) {
        self.head.borrow().next.clone().unwrap().borrow_mut().pre = Some(node.clone());

        node.borrow_mut().pre = Some(self.head.clone());
        node.borrow_mut().next = self.head.borrow().next.clone();

        self.head.borrow_mut().next = Some(node.clone());

        self.key_to_node.insert(node.borrow().key, node.clone());
    }
}

struct Node {
    key: i32,
    value: i32,

    pre: Option<Rc<RefCell<Node>>>,
    next: Option<Rc<RefCell<Node>>>,
}

impl Node {
    fn new(key: i32, value: i32) -> Rc<RefCell<Node>> {
        Rc::new(RefCell::new(Node {
            key,
            value,
            pre: None,
            next: None,
        }))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_lru() {
        let node = Node::new(1, 1);
        let a1 = node.clone();
        let a2 = node.clone();

        print!("a1 == a2: {}", Rc::ptr_eq(&a1, &a2))
    }

    #[test]
    fn test_lru_02() {
        let cache = &mut LRUCache::new(2);
        cache.put(1, 1);
        cache.put(2, 2);

        println!("get 1: {}", cache.get(1)); // 1
        cache.put(3, 3);
        println!("get 2: {}", cache.get(2)); // -1
        cache.put(4, 4);


        println!("get 1: {}", cache.get(1)); // -1
        println!("get 3: {}", cache.get(3)); // 3
        println!("get 4: {}", cache.get(4)); // 4

    }
}
