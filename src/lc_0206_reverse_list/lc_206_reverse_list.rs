pub struct Node {
    val: i32,
    next: Option<Box<Node>>,
}

impl Node {
    fn new(val: i32) -> Self {
        Node {
            val: val,
            next: None,
        }
    }

    fn new_with_next(val: i32, next: Option<Box<Node>>) -> Self {
        Node { val, next }
    }
}

pub fn reverse_list(head: Option<Box<Node>>) -> Option<Box<Node>> {
    let mut new_head = None;
    let mut iter_next_node = head;

    while let Some(mut cur_node) = iter_next_node {
        iter_next_node = cur_node.next.take();
        cur_node.next = new_head;
        new_head = Some(cur_node);
    }

    new_head
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_reverse_list() {
        let n5 = Some(Box::new(Node::new_with_next(5, None)));
        let n4 = Some(Box::new(Node::new_with_next(4, n5)));
        let n3 = Some(Box::new(Node::new_with_next(3, n4)));
        let n2 = Some(Box::new(Node::new_with_next(2, n3)));
        let n1 = Some(Box::new(Node::new_with_next(1, n2)));

        let head = reverse_list(n1);
        let mut node_tmp = head;
        while let Some(node) = node_tmp {
            println!("val: {}", node.val);
            node_tmp = node.next;
        }
    }
}
